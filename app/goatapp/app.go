package goatapp

import (
	"os"
	"runtime"
	"strings"

	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/plainmap"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/app/terminal"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/dependency/provider"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"
	"github.com/goatcms/goatcore/filesystem/json"
	"github.com/mitchellh/go-homedir"
)

// GoatApp is base app template
type GoatApp struct {
	app.AppHealthCheckers

	filespaces filespacesProvider
	injectors  []dependency.Injector
	ioContext  app.IOContext
	params     Params
	scopes     scopesProvider
}

// NewGoatApp create new app instance
func NewGoatApp(params Params) (a app.App, err error) {
	var (
		rootPath string
	)
	if runtime.GOOS == "windows" {
		rootPath = os.Getenv("SYSTEMDRIVE") + "\\"
	} else {
		rootPath = "/"
	}
	if params.Arguments == nil {
		params.Arguments = os.Args
	}
	if params.Name == "" {
		panic(ErrAppNameIsRequired)
	}
	if params.IO.Err == nil {
		params.IO.Err = gio.NewAppOutput(os.Stderr)
	}
	if params.IO.In == nil {
		params.IO.In = gio.NewAppInput(os.Stdin)
	}
	if params.IO.Out == nil {
		params.IO.Out = gio.NewAppOutput(os.Stdout)
	}
	if params.Terminal == nil {
		params.Terminal = terminal.NewTerminalManager()
	}
	if params.Version == nil {
		params.Version = NilVersion
	}
	// filespaces
	if params.Filespaces.Root == nil {
		if params.Filespaces.Root, err = diskfs.NewFilespace(rootPath); err != nil {
			return
		}
	}
	if params.Filespaces.CWD == nil {
		if params.Filespaces.CWD, err = diskfs.NewFilespace(CWDPath); err != nil {
			return
		}
	}
	if params.Filespaces.Home == nil {
		var homePath string
		if homePath, err = homedir.Dir(); err != nil {
			return
		}
		if err = params.Filespaces.Root.MkdirAll(homePath, filesystem.SafeDirPermissions); err != nil {
			return
		}
		if params.Filespaces.Home, err = params.Filespaces.Root.Filespace(homePath); err != nil {
			return
		}
	}
	if params.Filespaces.Tmp == nil {
		var (
			randID  = varutil.RandString(20, varutil.AlphaNumericBytes)
			tmpPath = rootPath + "/tmp/" + randID
		)
		if err = params.Filespaces.Root.MkdirAll(tmpPath, filesystem.SafeDirPermissions); err != nil {
			return
		}
		if params.Filespaces.Tmp, err = params.Filespaces.Root.Filespace(tmpPath); err != nil {
			return
		}
	}
	// filespace scopes
	fsCope := scope.NewScope(scope.Params{
		Tag: app.FilespaceTagName,
	})
	fsCope.SetValue(app.RootFilespace, params.Filespaces.Root)
	fsCope.SetValue(app.TmpFilespace, params.Filespaces.Tmp)
	fsCope.SetValue(app.HomeFilespace, params.Filespaces.Home)
	fsCope.SetValue(app.CurrentFilespace, params.Filespaces.CWD)
	params.Scopes.filespace = fsCope
	// args scope
	if params.Scopes.args, err = argscope.NewScope(params.Arguments, app.ArgsTagName); err != nil {
		return
	}
	// args scope - init envs
	if params.Env == "" {
		params.Env, _ = scope.GetString(params.Scopes.args, "env")
		if params.Env == "" {
			params.Env = app.DefaultEnv
		}
	}
	// config scopes
	if params.Scopes.config == nil {
		var (
			configPlainMap map[string]interface{}
			configMap      = make(map[string]interface{})
			configPath     = strings.Replace(ConfigFilePath, "{{env}}", params.Env, -1)
		)
		if params.Filespaces.CWD.IsFile(configPath) {
			if err = json.ReadJSON(params.Filespaces.CWD, configPath, &configMap); err != nil {
				return
			}
		}
		if configPlainMap, err = plainmap.RecursiveMapToPlainMap(configMap); err != nil {
			return
		}
		params.Scopes.config = scope.NewDataScope(varutil.ToMapInterfaceInterface(configPlainMap))
	}
	// filespace scopes
	fsScope := scope.NewScope(scope.Params{
		Tag: app.FilespaceTagName,
	})
	fsScope.SetValue(app.RootFilespace, params.Filespaces.Root)
	fsScope.SetValue(app.TmpFilespace, params.Filespaces.Tmp)
	fsScope.SetValue(app.HomeFilespace, params.Filespaces.Home)
	fsScope.SetValue(app.CurrentFilespace, params.Filespaces.CWD)
	params.Scopes.filespace = fsScope
	// app scope
	params.Scopes.App = scope.NewScope(scope.Params{
		Tag: app.AppTagName,
	})
	// Dependency provider
	if params.DP == nil {
		params.DP = provider.NewProvider(app.DependencyTagName)
	}
	params.DP.SetDefault(app.AppScope, params.Scopes.App)
	params.DP.SetDefault(app.ArgsScope, params.Scopes.args)
	params.DP.SetDefault(app.ConfigScope, params.Scopes.config)
	params.DP.SetDefault(app.ErrorService, params.IO.Err)
	params.DP.SetDefault(app.FilespaceScope, params.Scopes.filespace)
	params.DP.SetDefault(app.InputService, params.IO.In)
	params.DP.SetDefault(app.OutputService, params.IO.Out)
	params.DP.AddInjectors([]dependency.Injector{
		params.Scopes.App,
		scope.NewScopeInjector(app.ArgsTagName, params.Scopes.args),
		scope.NewScopeInjector(app.ConfigTagName, params.Scopes.config),
		scope.NewScopeInjector(app.FilespaceTagName, params.Scopes.filespace),
	})
	// helth
	if params.HealthCheckers == nil {
		params.HealthCheckers = newHealthCheckers()
	}
	// app object
	gapp := &GoatApp{
		AppHealthCheckers: params.HealthCheckers,

		filespaces: filespacesProvider{
			fs: params.Filespaces,
		},
		ioContext: gio.NewIOContext(params.Scopes.App, gio.NewIO(gio.IOParams{
			CWD: params.Filespaces.CWD,
			Err: params.IO.Err,
			In:  params.IO.In,
			Out: params.IO.Out,
		})),
		params: params,
		scopes: scopesProvider{
			scopes: params.Scopes,
		},
	}
	gapp.params.DP.SetDefault(app.AppService, app.App(gapp))
	return gapp, nil
}

// InjectTo inject application dependencies, arguments, config and filespaces to
func (gapp *GoatApp) InjectTo(obj interface{}) error {
	return gapp.params.DP.InjectTo(obj)
}

// Arguments return application arguments
func (gapp *GoatApp) Arguments() []string {
	return gapp.params.Arguments
}

// DependencyProvider return dependency provider
func (gapp *GoatApp) DependencyProvider() dependency.Provider {
	return gapp.params.DP
}

// Filespaces return filespaces
func (gapp *GoatApp) Filespaces() app.AppFilespaces {
	return gapp.filespaces
}

// IOContext return main IOContext fo application
func (gapp *GoatApp) IOContext() app.IOContext {
	return gapp.ioContext
}

// Name return app name
func (gapp *GoatApp) Name() string {
	return gapp.params.Name
}

// AppScope return app scope
func (gapp *GoatApp) Scopes() app.AppScopes {
	return gapp.scopes
}

// Terminal return app terminal manager
func (gapp *GoatApp) Terminal() app.TerminalManager {
	return gapp.params.Terminal
}

// Version return app version
func (gapp *GoatApp) Version() app.Version {
	return gapp.params.Version
}
