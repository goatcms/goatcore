package gtprovider

import (
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/fsloop"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Provider provide templates api
type Provider struct {
	fs           filesystem.Filespace
	helpersPath  string
	layoutPath   string
	viewPath     string
	extension    string
	baseMutex    sync.Mutex
	baseTemplate *template.Template
	layoutMutex  sync.Mutex
	layouts      map[string]*template.Template
	viewMutex    sync.Mutex
	views        map[string]*template.Template
	funcs        template.FuncMap
	isCached     bool
}

// NewProvider create Provider instance
func NewProvider(fs filesystem.Filespace, helpersPath, layoutPath, viewPath, extension string, funcs template.FuncMap, isCached bool) *Provider {
	return &Provider{
		fs:          fs,
		layoutPath:  layoutPath,
		helpersPath: helpersPath,
		viewPath:    viewPath,
		extension:   extension,
		layouts:     map[string]*template.Template{},
		views:       map[string]*template.Template{},
		funcs:       funcs,
		isCached:    isCached,
	}
}

// Base return base template (with loaded helpers)
func (provider *Provider) Base() (*template.Template, error) {
	if provider.baseTemplate != nil {
		return provider.baseTemplate, nil
	}
	return provider.base()
}

func (provider *Provider) base() (baseTemplate *template.Template, err error) {
	provider.baseMutex.Lock()
	defer provider.baseMutex.Unlock()
	if provider.baseTemplate != nil {
		return provider.baseTemplate, nil
	}
	baseTemplate = template.New("baseTemplate")
	baseTemplate.Funcs(provider.funcs)
	if !provider.fs.IsDir(provider.helpersPath) {
		if provider.isCached {
			provider.baseTemplate = baseTemplate
		}
		return baseTemplate, nil
	}
	templateLoader := NewTemplateLoader(baseTemplate)
	if err = fsloop.WalkFS(provider.fs, provider.helpersPath, func(path string, info os.FileInfo) (err error) {
		if !strings.HasSuffix(path, provider.extension) {
			return nil
		}
		return templateLoader.Load(provider.fs, path)
	}, nil); err != nil {
		return nil, err
	}
	if provider.isCached {
		provider.baseTemplate = baseTemplate
	}
	return baseTemplate, nil
}

// Layout return template for named layout (with loaded helpers and layout definitions)
func (provider *Provider) Layout(name string) (tmpl *template.Template, err error) {
	var (
		ok bool
	)
	if name == "" {
		name = goathtml.DefaultLayout
	}
	if tmpl, ok = provider.layouts[name]; ok {
		return tmpl, nil
	}
	return provider.layout(name)
}

func (provider *Provider) layout(name string) (layoutTemplate *template.Template, err error) {
	var (
		ok bool
	)
	provider.layoutMutex.Lock()
	defer provider.layoutMutex.Unlock()
	if layoutTemplate, ok = provider.layouts[name]; ok {
		return layoutTemplate, nil
	}
	if layoutTemplate, err = provider.Base(); err != nil {
		return nil, err
	}
	path := strings.Replace(provider.layoutPath, "{name}", name, 1)
	if !provider.fs.IsDir(path) {
		if provider.isCached {
			provider.layouts[name] = layoutTemplate
		}
		return layoutTemplate, nil
	}
	if layoutTemplate, err = layoutTemplate.Clone(); err != nil {
		return nil, err
	}
	templateLoader := NewTemplateLoader(layoutTemplate)
	if err = fsloop.WalkFS(provider.fs, path, func(path string, info os.FileInfo) (err error) {
		if !strings.HasSuffix(path, provider.extension) {
			return nil
		}
		return templateLoader.Load(provider.fs, path)
	}, nil); err != nil {
		return nil, err
	}
	if provider.isCached {
		provider.layouts[name] = layoutTemplate
	}
	return layoutTemplate, nil
}

// View return template for view by name. It contains selected layout definitions and helpers
func (provider *Provider) View(layoutName, viewName string) (tmpl *template.Template, err error) {
	var (
		ok  bool
		key string
	)
	if layoutName == "" {
		layoutName = goathtml.DefaultLayout
	}
	if viewName == "" {
		return nil, goaterr.Errorf("goathtml.Provider: A view name is required")
	}
	key = layoutName + ":" + viewName
	if tmpl, ok = provider.views[key]; ok {
		return tmpl, nil
	}
	return provider.view(layoutName, viewName, key)
}

func (provider *Provider) view(layoutName, viewName, key string) (viewTemplate *template.Template, err error) {
	var (
		ok   bool
		path string
	)
	provider.viewMutex.Lock()
	defer provider.viewMutex.Unlock()
	// check after lock
	if viewTemplate, ok = provider.views[key]; ok {
		return viewTemplate, nil
	}
	// create a new view
	if viewTemplate, err = provider.Layout(layoutName); err != nil {
		return nil, err
	}
	path = strings.Replace(provider.viewPath, "{name}", viewName, 1)
	if !provider.fs.IsDir(path) {
		if provider.isCached {
			provider.views[key] = viewTemplate
		}
		return viewTemplate, nil
	}
	if viewTemplate, err = viewTemplate.Clone(); err != nil {
		return nil, err
	}
	viewTemplate.Funcs(provider.funcs)
	templateLoader := NewTemplateLoader(viewTemplate)
	if err = fsloop.WalkFS(provider.fs, path, func(path string, info os.FileInfo) (err error) {
		if !strings.HasSuffix(path, provider.extension) {
			return nil
		}
		return templateLoader.Load(provider.fs, path)
	}, nil); err != nil {
		return nil, err
	}
	if provider.isCached {
		provider.views[key] = viewTemplate
	}
	return viewTemplate, nil
}
