package goatapp

import (
	"os"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
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
	io              app.IO
	ioContext       app.IOContext
	dp              dependency.Provider
	cwd             filesystem.Filespace
}

const (
	// ConfigJSONPath is path to main config file
	ConfigJSONPath = "/config/config_{{env}}.json"
)

// NewGoatApp create new app instance
func NewGoatApp(name, version, basePath string) (a app.App, err error) {
	var (
		in   = gio.NewAppInput(os.Stdin)
		out  = gio.NewAppOutput(os.Stdout)
		eout = gio.NewAppOutput(os.Stderr)
	)
	gapp := &GoatApp{
		name:    name,
		version: version,
	}

	if err = gapp.initEngineScope(); err != nil {
		return nil, err
	}
	if err = gapp.initArgsScope(); err != nil {
		return nil, err
	}
	if err = gapp.initFilespaceScope(basePath); err != nil {
		return nil, err
	}
	if err = gapp.initConfigScope(); err != nil {
		return nil, err
	}
	if err = gapp.initDependencyScope(); err != nil {
		return nil, err
	}
	if err = gapp.initAppScope(); err != nil {
		return nil, err
	}
	if err = gapp.initCommandScope(); err != nil {
		return nil, err
	}

	gapp.dp.SetDefault(app.EngineScope, gapp.engineScope)
	gapp.dp.SetDefault(app.ArgsScope, gapp.argsScope)
	gapp.dp.SetDefault(app.FilespaceScope, gapp.filespaceScope)
	gapp.dp.SetDefault(app.ConfigScope, gapp.configScope)
	gapp.dp.SetDefault(app.DependencyScope, gapp.dependencyScope)
	gapp.dp.SetDefault(app.AppScope, gapp.appScope)
	gapp.dp.SetDefault(app.CommandScope, gapp.commandScope)

	if gapp.io, err = gio.NewIO(in, out, eout, gapp.cwd); err != nil {
		return nil, err
	}
	if gapp.ioContext, err = gio.NewIOContext(gapp.appScope, gapp.io); err != nil {
		return nil, err
	}

	gapp.dp.SetDefault(app.InputService, gapp.io.In())
	gapp.dp.SetDefault(app.OutputService, gapp.io.Out())
	gapp.dp.SetDefault(app.ErrorService, gapp.io.Err())

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

	gapp.dp.SetDefault(app.AppService, app.App(gapp))
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

func (gapp *GoatApp) initFilespaceScope(path string) (err error) {
	var cwdi interface{}
	gapp.rootFilespace, err = diskfs.NewFilespace(path)
	if err != nil {
		return err
	}
	gapp.filespaceScope = scope.NewScope(app.FilespaceTagName)
	gapp.filespaceScope.Set(app.RootFilespace, gapp.rootFilespace)
	if err = gapp.rootFilespace.MkdirAll("tmp", 0766); err != nil {
		return err
	}
	tmpFilespace, err := gapp.rootFilespace.Filespace("tmp")
	if err != nil {
		return err
	}
	gapp.filespaceScope.Set(app.TmpFilespace, tmpFilespace)
	if cwdi, err = gapp.argsScope.Get("cwd"); err != nil {
		gapp.filespaceScope.Set(app.CurrentFilespace, gapp.rootFilespace)
	} else {
		var currentFS filesystem.Filespace
		if currentFS, err = diskfs.NewFilespace(cwdi.(string)); err != nil {
			return err
		}
		gapp.filespaceScope.Set(app.CurrentFilespace, currentFS)
	}
	return nil
}

func (gapp *GoatApp) initConfigScope() error {
	var (
		deps struct {
			Env string `argument:"?env"`
		}
		err error
	)
	if err = gapp.argsScope.InjectTo(&deps); err != nil {
		return err
	}
	if deps.Env == "" {
		deps.Env = app.DefaultEnv
	}
	fullmap := make(map[string]interface{})
	path := strings.Replace(ConfigJSONPath, "{{env}}", deps.Env, -1)
	if err = json.ReadJSON(gapp.rootFilespace, path, &fullmap); err != nil {
		return err
	}
	plainmap, err := plainmap.RecursiveMapToPlainMap(fullmap)
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

// RootFilespace return main filespace for application (current directory by default)
func (gapp *GoatApp) RootFilespace() filesystem.Filespace {
	return gapp.rootFilespace
}

// IO return main IO fo application
func (gapp *GoatApp) IO() app.IO {
	return gapp.io
}

// IOContext return main IOContext fo application
func (gapp *GoatApp) IOContext() app.IOContext {
	return gapp.ioContext
}
