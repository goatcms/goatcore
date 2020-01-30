package mockupapp

import (
	"bytes"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/dependency/provider"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// App is default mockup applicationo test code
type App struct {
	options   MockupOptions
	outBuf    bytes.Buffer
	errBuf    bytes.Buffer
	io        app.IO
	ioContext app.IOContext
}

// NewApp create a new App
func NewApp(options MockupOptions) (result *App, err error) {
	mapp := &App{
		options: options,
	}

	if mapp.options.RootFilespace == nil {
		if mapp.options.RootFilespace, err = memfs.NewFilespace(); err != nil {
			return nil, err
		}
	}
	if mapp.options.TMPFilespace == nil {
		rootfs := mapp.options.RootFilespace
		if err = rootfs.MkdirAll("tmp", 0766); err != nil {
			return nil, err
		}
		if mapp.options.TMPFilespace, err = rootfs.Filespace("tmp"); err != nil {
			return nil, err
		}
	}

	if mapp.options.Input == nil {
		mapp.options.Input = strings.NewReader("")
	}

	mapp.io = gio.NewIO(gio.IOParams{
		In:  gio.NewAppInput(mapp.options.Input),
		Out: gio.NewAppOutput(&mapp.outBuf),
		Err: gio.NewAppOutput(&mapp.errBuf),
		CWD: mapp.options.RootFilespace,
	})

	if err = goaterr.ToError(goaterr.AppendError(nil,
		mapp.initEngineScope(),
		mapp.initArgsScope(),
		mapp.initFilespaceScope(),
		mapp.initConfigScope(),
		mapp.initDependencyScope(),
		mapp.initAppScope(),
		mapp.initCommandScope(),
	)); err != nil {
		return nil, err
	}

	mapp.options.DP.SetDefault(app.EngineScope, mapp.options.EngineScope)
	mapp.options.DP.SetDefault(app.ArgsScope, mapp.options.ArgsScope)
	mapp.options.DP.SetDefault(app.FilespaceScope, mapp.options.FilespaceScope)
	mapp.options.DP.SetDefault(app.ConfigScope, mapp.options.ConfigScope)
	mapp.options.DP.SetDefault(app.AppScope, mapp.options.AppScope)
	mapp.options.DP.SetDefault(app.CommandScope, mapp.options.CommandScope)

	mapp.options.DP.SetDefault(app.InputService, mapp.io.In())
	mapp.options.DP.SetDefault(app.OutputService, mapp.io.Out())
	mapp.options.DP.SetDefault(app.ErrorService, mapp.io.Err())

	mapp.options.DP.AddInjectors([]dependency.Injector{
		mapp.options.CommandScope,
		mapp.options.AppScope,
		mapp.options.ConfigScope,
		mapp.options.FilespaceScope,
		mapp.options.ArgsScope,
		mapp.options.EngineScope,
	})

	mapp.ioContext = gio.NewIOContext(mapp.options.AppScope, mapp.io)

	mapp.options.DP.SetDefault(app.AppService, app.App(mapp))
	return mapp, nil
}

func (mapp *App) initEngineScope() error {
	if mapp.options.EngineScope != nil {
		return nil
	}
	mapp.options.EngineScope = scope.NewScope(scope.Params{
		Tag: app.EngineTagName,
	})
	mapp.options.EngineScope.Set(app.GoatVersion, app.GoatVersionValue)
	return nil
}

func (mapp *App) initArgsScope() (err error) {
	if mapp.options.ArgsScope != nil {
		return nil
	}
	var args []string
	if mapp.options.Args != nil {
		args = mapp.options.Args
	} else {
		args = []string{}
	}
	mapp.options.ArgsScope, err = argscope.NewScope(args, app.ArgsTagName)
	return err
}

func (mapp *App) initFilespaceScope() (err error) {
	var value interface{}
	if mapp.options.FilespaceScope == nil {
		mapp.options.FilespaceScope = scope.NewScope(scope.Params{
			Tag: app.FilespaceTagName,
		})
	}
	fsscope := mapp.options.FilespaceScope
	if value, _ = fsscope.Get(app.RootFilespace); value == nil {
		mapp.options.FilespaceScope.Set(app.RootFilespace, mapp.options.RootFilespace)
	}
	if value, _ = fsscope.Get(app.TmpFilespace); value == nil {
		mapp.options.FilespaceScope.Set(app.TmpFilespace, mapp.options.TMPFilespace)
	}
	if value, _ = fsscope.Get(app.CurrentFilespace); value == nil {
		mapp.options.FilespaceScope.Set(app.CurrentFilespace, mapp.io.CWD())
	}
	return nil
}

func (mapp *App) initConfigScope() error {
	if mapp.options.ConfigScope != nil {
		return nil
	}
	mapp.options.ConfigScope = scope.NewScope(scope.Params{
		Tag: app.ConfigTagName,
	})
	return nil
}

func (mapp *App) initCommandScope() error {
	if mapp.options.CommandScope != nil {
		return nil
	}
	mapp.options.CommandScope = scope.NewScope(scope.Params{
		Tag: app.CommandTagName,
	})
	return nil
}

func (mapp *App) initDependencyScope() error {
	if mapp.options.DP == nil {
		mapp.options.DP = provider.NewProvider(app.DependencyTagName)
	}
	return nil
}

func (mapp *App) initAppScope() error {
	if mapp.options.AppScope != nil {
		return nil
	}
	mapp.options.AppScope = scope.NewScope(scope.Params{
		Tag: app.AppTagName,
	})
	mapp.options.AppScope.Set(app.AppName, mapp.options.Name)
	mapp.options.AppScope.Set(app.AppVersion, mapp.options.Version)
	return nil
}

// Name return app name
func (mapp *App) Name() string {
	return mapp.options.Name
}

// Version return app version
func (mapp *App) Version() string {
	return mapp.options.Version
}

// Arguments return application arguments
func (mapp *App) Arguments() []string {
	return mapp.options.Args
}

// EngineScope return engine scope
func (mapp *App) EngineScope() app.Scope {
	return mapp.options.EngineScope
}

// ArgsScope return app scope
func (mapp *App) ArgsScope() app.Scope {
	return mapp.options.ArgsScope
}

// FilespaceScope return filespace scope
func (mapp *App) FilespaceScope() app.Scope {
	return mapp.options.FilespaceScope
}

// ConfigScope return config scope
func (mapp *App) ConfigScope() app.Scope {
	return mapp.options.ConfigScope
}

// AppScope return app scope
func (mapp *App) AppScope() app.Scope {
	return mapp.options.AppScope
}

// CommandScope return command scope
func (mapp *App) CommandScope() app.Scope {
	return mapp.options.CommandScope
}

// DependencyProvider return dependency provider
func (mapp *App) DependencyProvider() dependency.Provider {
	return mapp.options.DP
}

// RootFilespace return main filespace for application (current directory by default)
func (mapp *App) RootFilespace() filesystem.Filespace {
	return mapp.options.RootFilespace
}

// IO return main application Input/Output objcet
func (mapp *App) IO() app.IO {
	return mapp.io
}

// IOContext return application IO context
func (mapp *App) IOContext() app.IOContext {
	return mapp.ioContext
}

// OutputBuffer return output buffer
func (mapp *App) OutputBuffer() *bytes.Buffer {
	return &mapp.outBuf
}

// ErrorBuffer return error output buffer
func (mapp *App) ErrorBuffer() *bytes.Buffer {
	return &mapp.errBuf
}
