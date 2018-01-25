package ghprovider

import (
	"fmt"
	"html/template"
	"strings"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/fsloop"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

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
}

func NewProvider(fs filesystem.Filespace, helpersPath, layoutPath, viewPath, extension string, funcs template.FuncMap) *Provider {
	return &Provider{
		fs:          fs,
		layoutPath:  layoutPath,
		helpersPath: helpersPath,
		viewPath:    viewPath,
		extension:   extension,
		layouts:     map[string]*template.Template{},
		views:       map[string]*template.Template{},
		funcs:       funcs,
	}
}

func (provider *Provider) Base(eventScope app.EventScope) (*template.Template, error) {
	if provider.baseTemplate != nil {
		return provider.baseTemplate, nil
	}
	return provider.base(eventScope)
}

func (provider *Provider) base(eventScope app.EventScope) (*template.Template, error) {
	provider.baseMutex.Lock()
	defer provider.baseMutex.Unlock()
	if provider.baseTemplate != nil {
		return provider.baseTemplate, nil
	}
	baseTemplate := template.New("baseTemplate")
	baseTemplate.Funcs(provider.funcs)
	if !provider.fs.IsDir(provider.helpersPath) {
		return baseTemplate, nil
	}
	templateLoader := NewTemplateLoader(baseTemplate)
	loop := fsloop.NewLoop(&fsloop.LoopData{
		Filespace: provider.fs,
		FileFilter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, provider.extension)
		},
		OnFile:     templateLoader.Load,
		Producents: 1,
		Consumers:  1,
	}, eventScope)
	loop.Run(provider.helpersPath)
	loop.Wait()
	if len(loop.Errors()) != 0 {
		return nil, goaterr.NewErrors(loop.Errors())
	}
	provider.baseTemplate = baseTemplate
	return baseTemplate, nil
}

func (provider *Provider) Layout(name string, eventScope app.EventScope) (*template.Template, error) {
	if name == "" {
		name = goathtml.DefaultLayout
	}
	tmpl, ok := provider.layouts[name]
	if ok {
		return tmpl, nil
	}
	return provider.layout(name, eventScope)
}

func (provider *Provider) layout(name string, eventScope app.EventScope) (*template.Template, error) {
	provider.layoutMutex.Lock()
	defer provider.layoutMutex.Unlock()
	tmpl, ok := provider.layouts[name]
	if ok {
		return tmpl, nil
	}
	layoutTemplate, err := provider.Base(eventScope)
	if err != nil {
		return nil, err
	}
	path := strings.Replace(provider.layoutPath, "{name}", name, 1)
	if !provider.fs.IsDir(path) {
		return layoutTemplate, nil
	}
	templateLoader := NewTemplateLoader(layoutTemplate)
	loop := fsloop.NewLoop(&fsloop.LoopData{
		Filespace: provider.fs,
		FileFilter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, provider.extension)
		},
		OnFile:     templateLoader.Load,
		Producents: 1,
		Consumers:  1,
	}, eventScope)
	loop.Run(path)
	loop.Wait()
	if len(loop.Errors()) != 0 {
		return nil, goaterr.NewErrors(loop.Errors())
	}
	provider.layouts[name] = layoutTemplate
	return layoutTemplate, nil
}

func (provider *Provider) View(layoutName, viewName string, eventScope app.EventScope) (*template.Template, error) {
	if layoutName == "" {
		layoutName = goathtml.DefaultLayout
	}
	if viewName == "" {
		return nil, fmt.Errorf("goathtml.Provider: A view name is required")
	}
	key := layoutName + ":" + viewName
	// check without lock (preformence feature)
	tmpl, ok := provider.views[key]
	if ok {
		return tmpl, nil
	}
	return provider.view(layoutName, viewName, key, eventScope)
}

func (provider *Provider) view(layoutName, viewName, key string, eventScope app.EventScope) (*template.Template, error) {
	provider.viewMutex.Lock()
	defer provider.viewMutex.Unlock()
	// check after lock
	tmpl, ok := provider.views[key]
	if ok {
		return tmpl, nil
	}
	// create a new view
	layoutTemplate, err := provider.Layout(layoutName, eventScope)
	if err != nil {
		return nil, err
	}
	viewTemplate, err := layoutTemplate.Clone()
	viewTemplate.Funcs(provider.funcs)
	if err != nil {
		return nil, err
	}
	templateLoader := NewTemplateLoader(viewTemplate)
	loop := fsloop.NewLoop(&fsloop.LoopData{
		Filespace: provider.fs,
		FileFilter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, provider.extension)
		},
		OnFile:     templateLoader.Load,
		Consumers:  1,
		Producents: 1,
	}, eventScope)
	path := strings.Replace(provider.viewPath, "{name}", viewName, 1)
	loop.Run(path)
	loop.Wait()
	if len(loop.Errors()) != 0 {
		return nil, goaterr.NewErrors(loop.Errors())
	}
	tmpl = templateLoader.Template()
	provider.views[key] = tmpl
	return tmpl, nil
}
