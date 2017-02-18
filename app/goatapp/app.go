package goatapp

import (
	"os"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/dependency/provider"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"
	"github.com/goatcms/goatcore/filesystem/json"
	"github.com/goatcms/goatcore/varutil/plainmap"
)

// GoatApp is base app template
type GoatApp struct {
	name    string
	version string

	rootFilespace filesystem.Filespace

	engineScope     app.Scope
	argsScope       app.Scope
	filespaceScope  app.Scope
	configScope     app.Scope
	dependencyScope app.Scope
	appScope        app.Scope
	commandScope    app.Scope

	dp dependency.Provider
}

const (
	// ConfigJSONPath is path to main config file
	ConfigJSONPath = "/config/config.json"
)

// NewGoatApp create new app instance
func NewGoatApp(name, version, basePath string) (app.App, error) {
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

	gapp.dp.SetDefault(app.EngineScope, gapp.engineScope)
	gapp.dp.SetDefault(app.ArgsScope, gapp.argsScope)
	gapp.dp.SetDefault(app.FilespaceScope, gapp.filespaceScope)
	gapp.dp.SetDefault(app.ConfigScope, gapp.configScope)
	gapp.dp.SetDefault(app.DependencyScope, gapp.dependencyScope)
	gapp.dp.SetDefault(app.AppScope, gapp.appScope)
	gapp.dp.SetDefault(app.CommandScope, gapp.commandScope)

	gapp.dp.AddInjectors([]dependency.Injector{
		gapp.commandScope,
		gapp.appScope,
		// gapp.dependencyScope, <- it is wraper for dependency injection and musn't
		// contains recursive injection
		gapp.configScope,
		gapp.filespaceScope,
		gapp.argsScope,
		gapp.engineScope,
	})

	return gapp, nil
}

func (gapp *GoatApp) initEngineScope() error {
	gapp.engineScope = scope.NewScope(app.EngineTagName)
	gapp.engineScope.Set(app.GoatVersion, app.GoatVersionValue)
	return nil
}

func (gapp *GoatApp) initArgsScope() error {
	var err error
	gapp.argsScope, err = argscope.NewScope(os.Args, app.ArgsTagName)
	return err
}

func (gapp *GoatApp) initFilespaceScope(path string) error {
	var err error
	gapp.rootFilespace, err = diskfs.NewFilespace(path)
	if err != nil {
		return err
	}
	gapp.filespaceScope = scope.NewScope(app.FilespaceTagName)
	gapp.filespaceScope.Set(app.RootFilespace, gapp.rootFilespace)
	tmpFilespace, err := gapp.rootFilespace.Filespace("tmp")
	if err != nil {
		return err
	}
	gapp.filespaceScope.Set(app.TmpFilespace, tmpFilespace)
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
	gapp.dp = provider.NewProvider(app.DependencyTagName)
	gapp.dependencyScope = NewDependencyScope(gapp.dp)
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

// DependencyProvider return dependency provider
func (gapp *GoatApp) DependencyProvider() dependency.Provider {
	return gapp.dp
}

func (gapp *GoatApp) RootFilespace() filesystem.Filespace {
	return gapp.rootFilespace
}
