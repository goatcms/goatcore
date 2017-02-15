package ghprovider

import (
	"fmt"
	"html/template"
	"strings"
	"sync"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/loop"
	"github.com/goatcms/goat-core/goathtml"
)

type Provider struct {
	fs          filesystem.Filespace
	layoutPath  string
	viewPath    string
	layoutMutex sync.Mutex
	layouts     map[string]*template.Template
	viewMutex   sync.Mutex
	views       map[string]*template.Template
	funcs       template.FuncMap
}

func NewProvider(fs filesystem.Filespace, layoutPath, viewPath string, funcs template.FuncMap) *Provider {
	return &Provider{
		fs:         fs,
		layoutPath: layoutPath,
		viewPath:   viewPath,
		layouts:    map[string]*template.Template{},
		views:      map[string]*template.Template{},
		funcs:      funcs,
	}
}

func (provider *Provider) Layout(name string, eventScope app.EventScope) (*template.Template, []error) {
	if name == "" {
		name = goathtml.DefaultLayout
	}
	tmpl, ok := provider.layouts[name]
	if ok {
		return tmpl, nil
	}
	return provider.layout(name, eventScope)
}

func (provider *Provider) layout(name string, eventScope app.EventScope) (*template.Template, []error) {
	provider.layoutMutex.Lock()
	defer provider.layoutMutex.Unlock()
	tmpl, ok := provider.layouts[name]
	if ok {
		return tmpl, nil
	}
	layoutTemplate := template.New(name)
	layoutTemplate.Funcs(provider.funcs)
	templateLoader := NewTemplateLoader(layoutTemplate)
	loop := loop.NewLoop(&loop.LoopData{
		Filespace: provider.fs,
		FileFilter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, goathtml.FileExtension)
		},
		OnFile: templateLoader.Load,
	}, eventScope)
	path := strings.Replace(provider.layoutPath, "{name}", name, 1)
	loop.Run(path)
	loop.Wait()
	if len(loop.Errors()) != 0 {
		return nil, loop.Errors()
	}
	provider.layouts[name] = layoutTemplate
	return layoutTemplate, nil
}

func (provider *Provider) View(layoutName, viewName string, eventScope app.EventScope) (*template.Template, []error) {
	if layoutName == "" {
		layoutName = goathtml.DefaultLayout
	}
	if viewName == "" {
		return nil, []error{fmt.Errorf("goathtml.Provider: A view name is required")}
	}
	key := layoutName + ":" + viewName
	// check without lock (preformence feature)
	tmpl, ok := provider.views[key]
	if ok {
		return tmpl, nil
	}
	return provider.view(layoutName, viewName, key, eventScope)
}

func (provider *Provider) view(layoutName, viewName, key string, eventScope app.EventScope) (*template.Template, []error) {
	provider.viewMutex.Lock()
	defer provider.viewMutex.Unlock()
	// check after lock
	tmpl, ok := provider.views[key]
	if ok {
		return tmpl, nil
	}
	// create a new view
	layoutTemplate, errs := provider.Layout(layoutName, eventScope)
	if errs != nil {
		return nil, errs
	}
	viewTemplate, err := layoutTemplate.Clone()
	viewTemplate.Funcs(provider.funcs)
	if err != nil {
		return nil, []error{err}
	}
	templateLoader := NewTemplateLoader(viewTemplate)
	loop := loop.NewLoop(&loop.LoopData{
		Filespace: provider.fs,
		FileFilter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, goathtml.FileExtension)
		},
		OnFile: templateLoader.Load,
	}, eventScope)
	path := strings.Replace(provider.viewPath, "{name}", viewName, 1)
	loop.Run(path)
	loop.Wait()
	if len(loop.Errors()) != 0 {
		return nil, loop.Errors()
	}
	tmpl = templateLoader.Template()
	provider.views[key] = tmpl
	return tmpl, nil
}
