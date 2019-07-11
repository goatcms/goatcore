package mockupapp

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/dependency/provider"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

// MockupApp is default mockup applicationo test code
// It is used t
type MockupApp struct {
	options MockupOptions
}

// NewApp create a new MockupApp
func NewApp(options MockupOptions) (app.App, error) {
	mapp := &MockupApp{
		options: options,
	}
	if err := mapp.initEngineScope(); err != nil {
		return nil, err
	}
	if err := mapp.initArgsScope(); err != nil {
		return nil, err
	}
	if err := mapp.initFilespaceScope(); err != nil {
		return nil, err
	}
	if err := mapp.initConfigScope(); err != nil {
		return nil, err
	}
	if err := mapp.initDependencyScope(); err != nil {
		return nil, err
	}
	if err := mapp.initAppScope(); err != nil {
		return nil, err
	}
	if err := mapp.initCommandScope(); err != nil {
		return nil, err
	}
	mapp.options.DP.SetDefault(app.EngineScope, mapp.options.EngineScope)
	mapp.options.DP.SetDefault(app.ArgsScope, mapp.options.ArgsScope)
	mapp.options.DP.SetDefault(app.FilespaceScope, mapp.options.FilespaceScope)
	mapp.options.DP.SetDefault(app.ConfigScope, mapp.options.ConfigScope)
	mapp.options.DP.SetDefault(app.DependencyScope, mapp.options.DependencyScope)
	mapp.options.DP.SetDefault(app.AppScope, mapp.options.AppScope)
	mapp.options.DP.SetDefault(app.CommandScope, mapp.options.CommandScope)
	mapp.options.DP.SetDefault(app.InputService, mapp.options.Input)
	mapp.options.DP.SetDefault(app.OutputService, mapp.options.Output)
	mapp.options.DP.AddInjectors([]dependency.Injector{
		mapp.options.CommandScope,
		mapp.options.AppScope,
		mapp.options.ConfigScope,
		mapp.options.FilespaceScope,
		mapp.options.ArgsScope,
		mapp.options.EngineScope,
	})
	mapp.options.DP.SetDefault(app.AppService, app.App(mapp))
	return mapp, nil
}

func (mapp *MockupApp) initEngineScope() error {
	if mapp.options.EngineScope != nil {
		return nil
	}
	mapp.options.EngineScope = scope.NewScope(app.EngineTagName)
	mapp.options.EngineScope.Set(app.GoatVersion, app.GoatVersionValue)
	return nil
}

func (mapp *MockupApp) initArgsScope() (err error) {
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

func (mapp *MockupApp) initFilespaceScope() (err error) {
	if err = mapp.initRootFilespace(); err != nil {
		return err
	}
	if err = mapp.initTMPFilespace(); err != nil {
		return err
	}
	if mapp.options.FilespaceScope != nil {
		return nil
	}
	mapp.options.FilespaceScope = scope.NewScope(app.FilespaceTagName)
	mapp.options.FilespaceScope.Set(app.RootFilespace, mapp.options.RootFilespace)
	mapp.options.FilespaceScope.Set(app.TmpFilespace, mapp.options.TMPFilespace)
	return nil
}

func (mapp *MockupApp) initRootFilespace() (err error) {
	if mapp.options.RootFilespace != nil {
		return nil
	}
	if mapp.options.RootFilespace, err = memfs.NewFilespace(); err != nil {
		return err
	}
	return nil
}

func (mapp *MockupApp) initTMPFilespace() (err error) {
	if mapp.options.TMPFilespace != nil {
		return nil
	}
	if err = mapp.options.RootFilespace.MkdirAll("tmp", 0766); err != nil {
		return err
	}
	if mapp.options.TMPFilespace, err = mapp.options.RootFilespace.Filespace("tmp"); err != nil {
		return err
	}
	return nil
}

func (mapp *MockupApp) initConfigScope() error {
	if mapp.options.ConfigScope != nil {
		return nil
	}
	mapp.options.ConfigScope = scope.NewScope(app.ConfigTagName)
	return nil
}

func (mapp *MockupApp) initCommandScope() error {
	if mapp.options.CommandScope != nil {
		return nil
	}
	mapp.options.CommandScope = scope.NewScope(app.CommandTagName)
	return nil
}

func (mapp *MockupApp) initDependencyScope() error {
	if mapp.options.DependencyScope != nil {
		return nil
	}
	mapp.options.DP = provider.NewProvider(app.DependencyTagName)
	mapp.options.DependencyScope = NewDependencyScope(mapp.options.DP)
	return nil
}

func (mapp *MockupApp) initAppScope() error {
	if mapp.options.AppScope != nil {
		return nil
	}
	mapp.options.AppScope = scope.NewScope(app.AppTagName)
	mapp.options.AppScope.Set(app.AppName, mapp.options.Name)
	mapp.options.AppScope.Set(app.AppVersion, mapp.options.Version)
	return nil
}

// Name return app name
func (mapp *MockupApp) Name() string {
	return mapp.options.Name
}

// Version return app version
func (mapp *MockupApp) Version() string {
	return mapp.options.Version
}

// EngineScope return engine scope
func (mapp *MockupApp) EngineScope() app.Scope {
	return mapp.options.EngineScope
}

// ArgsScope return app scope
func (mapp *MockupApp) ArgsScope() app.Scope {
	return mapp.options.ArgsScope
}

// FilespaceScope return filespace scope
func (mapp *MockupApp) FilespaceScope() app.Scope {
	return mapp.options.FilespaceScope
}

// ConfigScope return config scope
func (mapp *MockupApp) ConfigScope() app.Scope {
	return mapp.options.ConfigScope
}

// DependencyScope return dependency scope
func (mapp *MockupApp) DependencyScope() app.Scope {
	return mapp.options.DependencyScope
}

// AppScope return app scope
func (mapp *MockupApp) AppScope() app.Scope {
	return mapp.options.AppScope
}

// CommandScope return command scope
func (mapp *MockupApp) CommandScope() app.Scope {
	return mapp.options.CommandScope
}

// DependencyProvider return dependency provider
func (mapp *MockupApp) DependencyProvider() dependency.Provider {
	return mapp.options.DP
}

// RootFilespace return main filespace for application (current directory by default)
func (mapp *MockupApp) RootFilespace() filesystem.Filespace {
	return mapp.options.RootFilespace
}
