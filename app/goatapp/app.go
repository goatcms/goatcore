package goatapp

import (
	"os"
	"runtime"
	"strings"

	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/varutil/plainmap"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/dependency/provider"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"
	"github.com/goatcms/goatcore/filesystem/json"
	"github.com/mitchellh/go-homedir"
)

// GoatApp is base app template
type GoatApp struct {
	name    string
	version string

	arguments []string

	rootFilespace    filesystem.Filespace
	currentFilespace filesystem.Filespace
	homeFilespace    filesystem.Filespace

	engineScope    app.Scope
	argsScope      app.Scope
	filespaceScope app.Scope
	configScope    app.Scope
	appScope       app.Scope
	commandScope   app.Scope
	io             app.IO
	ioContext      app.IOContext
	dp             dependency.Provider
}

const (
	// ConfigJSONPath is path to main config file
	ConfigJSONPath = "/config/config_{{env}}.json"
)

// NewGoatApp create new app instance
func NewGoatApp(name, version, defaultCWDPath string) (a app.App, err error) {
	var (
		in   = gio.NewAppInput(os.Stdin)
		out  = gio.NewAppOutput(os.Stdout)
		eout = gio.NewAppOutput(os.Stderr)
	)
	gapp := &GoatApp{
		name:      name,
		version:   version,
		arguments: os.Args,
	}

	if err = gapp.initEngineScope(); err != nil {
		return nil, err
	}
	if err = gapp.initArgsScope(); err != nil {
		return nil, err
	}
	if err = gapp.initFilespaceScope(defaultCWDPath); err != nil {
		return nil, err
	}
	if err = gapp.initConfigScope(); err != nil {
		return nil, err
	}
	if err = gapp.initAppScope(); err != nil {
		return nil, err
	}
	if err = gapp.initCommandScope(); err != nil {
		return nil, err
	}

	gapp.dp = provider.NewProvider(app.DependencyTagName)

	gapp.dp.SetDefault(app.EngineScope, gapp.engineScope)
	gapp.dp.SetDefault(app.ArgsScope, gapp.argsScope)
	gapp.dp.SetDefault(app.FilespaceScope, gapp.filespaceScope)
	gapp.dp.SetDefault(app.ConfigScope, gapp.configScope)
	gapp.dp.SetDefault(app.AppScope, gapp.appScope)
	gapp.dp.SetDefault(app.CommandScope, gapp.commandScope)

	gapp.io = gio.NewIO(gio.IOParams{
		In:  in,
		Out: out,
		Err: eout,
		CWD: gapp.currentFilespace,
	})
	gapp.ioContext = gio.NewIOContext(gapp.appScope, gapp.io)

	gapp.dp.SetDefault(app.InputService, gapp.io.In())
	gapp.dp.SetDefault(app.OutputService, gapp.io.Out())
	gapp.dp.SetDefault(app.ErrorService, gapp.io.Err())

	gapp.dp.AddInjectors([]dependency.Injector{
		gapp.commandScope,
		gapp.appScope,
		gapp.configScope,
		gapp.filespaceScope,
		gapp.argsScope,
		gapp.engineScope,
	})

	gapp.dp.SetDefault(app.AppService, app.App(gapp))
	return gapp, nil
}

func (gapp *GoatApp) initEngineScope() error {
	gapp.engineScope = scope.NewScope(scope.Params{
		Tag: app.EngineTagName,
	})
	gapp.engineScope.SetValue(app.GoatVersion, app.GoatVersionValue)
	return nil
}

func (gapp *GoatApp) initArgsScope() error {
	var err error
	gapp.argsScope, err = argscope.NewScope(gapp.arguments, app.ArgsTagName)
	return err
}

func (gapp *GoatApp) initFilespaceScope(defaultCWDPath string) (err error) {
	var (
		rootPath  string
		homePath  string
		cwdPath   string
		tmpFS     filesystem.Filespace
		homeFS    filesystem.Filespace
		currentFS filesystem.Filespace
	)
	// create scope
	gapp.filespaceScope = scope.NewScope(scope.Params{
		Tag: app.FilespaceTagName,
	})
	// root filespace
	if runtime.GOOS == "windows" {
		rootPath = os.Getenv("SYSTEMDRIVE") + "\\"
	} else {
		rootPath = "/"
	}
	if gapp.rootFilespace, err = diskfs.NewFilespace(rootPath); err != nil {
		return err
	}
	gapp.filespaceScope.SetValue(app.RootFilespace, gapp.rootFilespace)
	// tmp filespace
	if err = gapp.rootFilespace.MkdirAll("tmp", 0766); err != nil {
		return err
	}
	if tmpFS, err = gapp.rootFilespace.Filespace("tmp"); err != nil {
		return err
	}
	gapp.filespaceScope.SetValue(app.TmpFilespace, tmpFS)
	// home filespace
	if homePath, err = homedir.Dir(); err != nil {
		return err
	}
	if err = gapp.rootFilespace.MkdirAll(homePath, filesystem.SafeDirPermissions); err != nil {
		return err
	}
	if homeFS, err = gapp.rootFilespace.Filespace(homePath); err != nil {
		return err
	}
	gapp.filespaceScope.SetValue(app.HomeFilespace, homeFS)
	// CWD (Current Working Directory)
	if cwdPath, err = scope.GetString(gapp.argsScope, "cwd"); cwdPath == "" || err != nil {
		cwdPath = defaultCWDPath
	}
	if !disk.IsDir(cwdPath) {
		return goaterr.Errorf("CWD (Current Working directory) path (%s) is not a directory", cwdPath)
	}
	if currentFS, err = diskfs.NewFilespace(cwdPath); err != nil {
		return err
	}
	gapp.filespaceScope.SetValue(app.CurrentFilespace, currentFS)
	gapp.currentFilespace = currentFS
	return nil
}

func (gapp *GoatApp) initConfigScope() error {
	var (
		deps struct {
			Env string `argument:"?env"`
		}
		pmap map[string]interface{}
		err  error
	)
	if err = gapp.argsScope.InjectTo(&deps); err != nil {
		return err
	}
	if deps.Env == "" {
		deps.Env = app.DefaultEnv
	}
	fullmap := make(map[string]interface{})
	path := strings.Replace(ConfigJSONPath, "{{env}}", deps.Env, -1)
	if gapp.currentFilespace.IsFile(path) {
		if err = json.ReadJSON(gapp.currentFilespace, path, &fullmap); err != nil {
			return err
		}
	}
	if pmap, err = plainmap.RecursiveMapToPlainMap(fullmap); err != nil {
		return err
	}
	ds := scope.NewDataScope(varutil.ToMapInterfaceInterface(pmap))
	gapp.configScope = &scope.Scope{
		EventScope: scope.NewEventScope(),
		DataScope:  ds,
		Injector:   scope.NewScopeInjector(app.ConfigTagName, ds),
	}
	return nil
}

func (gapp *GoatApp) initCommandScope() error {
	gapp.commandScope = scope.NewScope(scope.Params{
		Tag: app.CommandTagName,
	})
	return nil
}

func (gapp *GoatApp) initAppScope() error {
	gapp.appScope = scope.NewScope(scope.Params{
		Tag: app.AppTagName,
	})
	gapp.appScope.SetValue(app.AppName, gapp.name)
	gapp.appScope.SetValue(app.AppVersion, gapp.version)
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

// Arguments return application arguments
func (gapp *GoatApp) Arguments() []string {
	return gapp.arguments
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

// RootFilespace return root filespace for application ('/' directory by default)
func (gapp *GoatApp) RootFilespace() filesystem.Filespace {
	return gapp.rootFilespace
}

// HomeFilespace return user home directory filespace (like '/Users/username')
func (gapp *GoatApp) HomeFilespace() filesystem.Filespace {
	return gapp.homeFilespace
}

// IO return main IO fo application
func (gapp *GoatApp) IO() app.IO {
	return gapp.io
}

// IOContext return main IOContext fo application
func (gapp *GoatApp) IOContext() app.IOContext {
	return gapp.ioContext
}
