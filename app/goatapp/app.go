package goatapp

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/app/args"
	"github.com/goatcms/goat-core/app/injector"
	"github.com/goatcms/goat-core/app/scope"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/filesystem/filespace/diskfs"
	"github.com/goatcms/goat-core/filesystem/json"
	"github.com/goatcms/goat-core/varutil/plainmap"
)

// GoatApp is base app template
type GoatApp struct {
	name    string
	version string

	rootFilespace filesystem.Filespace

	globalScope     app.Scope
	engineScope     app.Scope
	argsScope       app.Scope
	filespaceScope  app.Scope
	configScope     app.Scope
	dependencyScope app.Scope
	appScope        app.Scope
	commandScope    app.Scope
}

const (
	// ConfigJSONPath is path to main config file
	ConfigJSONPath = "/config/config.json"
)

// NewGoatApp create new app instance
func NewGoatApp(name, version, basePath string) (*GoatApp, error) {
	gapp := &GoatApp{
		name:    name,
		version: version,
	}

	if err := gapp.initEngineScope(); err != nil {
		return nil, err
	}
	if err := gapp.initArgsScope(); err != nil {
		return nil, err
	}
	if err := gapp.initFilespaceScope(basePath); err != nil {
		return nil, err
	}
	if err := gapp.initConfigScope(); err != nil {
		return nil, err
	}
	if err := gapp.initDependencyScope(); err != nil {
		return nil, err
	}
	if err := gapp.initAppScope(); err != nil {
		return nil, err
	}
	if err := gapp.initCommandScope(); err != nil {
		return nil, err
	}

	dataScope := &scope.DataScope{
		Data: make(map[string]interface{}),
	}
	gapp.globalScope = scope.Scope{
		EventScope: scope.NewEventScope(),
		DataScope:  dataScope,
		Injector: injector.NewMultiInjector([]app.Injector{
			gapp.commandScope,
			gapp.appScope,
			gapp.dependencyScope,
			gapp.configScope,
			gapp.filespaceScope,
			gapp.argsScope,
			gapp.engineScope,
			dataScope.Injector(app.GlobalTagName),
		}),
	}

	gapp.globalScope.Set(app.EngineScope, gapp.engineScope)
	gapp.globalScope.Set(app.ArgsScope, gapp.argsScope)
	gapp.globalScope.Set(app.FilespaceScope, gapp.filespaceScope)
	gapp.globalScope.Set(app.ConfigScope, gapp.configScope)
	gapp.globalScope.Set(app.DependencyScope, gapp.dependencyScope)
	gapp.globalScope.Set(app.AppScope, gapp.appScope)
	gapp.globalScope.Set(app.CommandScope, gapp.commandScope)

	return gapp, nil
}

func (gapp *GoatApp) initEngineScope() error {
	gapp.argsScope = scope.Scope{
		EventScope: scope.NewEventScope(),
		DataScope:  scope.NewDataScope(map[string]interface{}{}),
		Injector:   args.NewInjector(app.EngineTagName),
	}
	gapp.engineScope.Set(app.GoatVersion, app.GoatVersionValue)
	return nil
}

func (gapp *GoatApp) initArgsScope() error {
	gapp.argsScope = scope.Scope{
		EventScope: scope.NewEventScope(),
		DataScope:  scope.NewDataScope(map[string]interface{}{}),
		Injector:   args.NewInjector(app.ArgsTagName),
	}
	return nil
}

func (gapp *GoatApp) initFilespaceScope(path string) error {
	var err error
	gapp.rootFilespace, err = diskfs.NewFilespace(path)
	if err != nil {
		return err
	}
	gapp.filespaceScope = scope.NewScope(app.FilespaceTagName)
	gapp.filespaceScope.Set(app.RootFilespace, gapp.rootFilespace)
	return nil
}

func (gapp *GoatApp) initConfigScope() error {
	var fullmap map[string]interface{}
	json.ReadJSON(gapp.rootFilespace, ConfigJSONPath, fullmap)
	plainmap, err := plainmap.ToPlainMap(fullmap)
	if err != nil {
		return err
	}
	ds := &scope.DataScope{
		Data: plainmap,
	}
	gapp.configScope = scope.Scope{
		EventScope: scope.NewEventScope(),
		DataScope:  ds,
		Injector:   ds.Injector(app.ConfigTagName),
	}
	return nil
}

func (gapp *GoatApp) initCommandScope() error {
	gapp.commandScope = scope.NewScope(app.CommandTagName)
	return nil
}

func (gapp *GoatApp) initDependencyScope() error {
	gapp.dependencyScope = NewDependencyScope(app.DependencyTagName)
	return nil
}

func (gapp *GoatApp) initAppScope() error {
	gapp.appScope = scope.NewScope(app.AppTagName)
	gapp.appScope.Set(app.AppName, gapp.name)
	gapp.appScope.Set(app.AppVersion, gapp.version)
	return nil
}

// Name return app name
func (gapp *GoatApp) Name() string {
	return gapp.name
}

// Version return app version
func (gapp *GoatApp) Version() string {
	return gapp.version
}

// GlobalScope return global scope
func (gapp *GoatApp) GlobalScope() app.Scope {
	return gapp.engineScope
}

// EngineScope return engine scope
func (gapp *GoatApp) EngineScope() app.Scope {
	return gapp.engineScope
}

// ArgsScope return app scope
func (gapp *GoatApp) ArgsScope() app.Scope {
	return gapp.argsScope
}

// FilespaceScope return filespace scope
func (gapp *GoatApp) FilespaceScope() app.Scope {
	return gapp.filespaceScope
}

// ConfigScope return config scope
func (gapp *GoatApp) ConfigScope() app.Scope {
	return gapp.configScope
}

// DependencyScope return dependency scope
func (gapp *GoatApp) DependencyScope() app.Scope {
	return gapp.dependencyScope
}

// AppScope return app scope
func (gapp *GoatApp) AppScope() app.Scope {
	return gapp.appScope
}

// CommandScope return command scope
func (gapp *GoatApp) CommandScope() app.Scope {
	return gapp.commandScope
}
