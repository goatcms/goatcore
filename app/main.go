package app

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/filesystem"
)

const (
	// KillEvent is kill action for current event
	KillEvent = iota
	// ErrorEvent is action for a error
	ErrorEvent = iota
	// CommitEvent is a action run when data is persist
	CommitEvent = iota
	// RollbackEvent is a action run when data is restore
	RollbackEvent = iota
	// BeforeCloseEvent is a action run during application close
	BeforeCloseEvent = iota

	// Error is key for error value
	Error = "error"

	// CLI is cli dependency name for engine scope
	CLI = "cli"

	// AppName is key to get application name (from app scope)
	AppName = "AppName"
	// AppVersion is key to get application version (from app scope)
	AppVersion = "AppVersion"
	// AppWelcome is key to get welcome message (from app scope)
	AppWelcome = "AppWelcome"
	// AppCompany is key to get owner company name (from app scope)
	AppCompany = "AppCompany"

	// GoatVersionValue represent goat app engine version
	GoatVersionValue = -0.001
	// GoatVersion is key to get goat version
	GoatVersion = "GoatVersion"

	// GlobalTagName is a name for global vars / consts injection
	GlobalTagName = "global"
	// EngineTagName is a name for app vars / const injection
	EngineTagName = "engine"
	// ArgsTagName is a name for argument injection
	ArgsTagName = "arg"
	// FilespaceTagName is a name for filepsace injection
	FilespaceTagName = "filespace"
	// ConfigTagName is a name for config injection
	ConfigTagName = "config"
	// DependencyTagName is a name for dependency injection
	DependencyTagName = "dependency"
	// AppTagName is a name for application data injection
	AppTagName = "app"
	// CommandTagName is a name for command injection
	CommandTagName = "command"

	// EngineScope is an engine scope
	EngineScope = "EngineScope"
	// ArgsScope is an arguments scope
	ArgsScope = "ArgScope"
	// FilespaceScope is a filespace scope
	FilespaceScope = "FilespaceScope"
	// ConfigScope is a config scope
	ConfigScope = "ConfigScope"
	// DependencyScope is a config scope
	DependencyScope = "DependencyScope"
	// AppScope is a application scope
	AppScope = "AppScope"
	// CommandScope is a command scope
	CommandScope = "CommandScope"
	// CommandScope is a command scope
	GlobalScope = "GlobalScope"

	//RootFilespace is key for root filesystem.Filespace
	RootFilespace = "root"

	// DefaultDurationValue is a default value for undefined env, configs etc
	DefaultDurationValue = 0
	// DefaultBoolValue is a default value for undefined env, configs etc
	DefaultBoolValue = false
	// DefaultStringValue is a default value for undefined env, configs etc
	DefaultStringValue = ""
	// DefaultFloat64Value is a default value for undefined env, configs etc
	DefaultFloat64Value = 0
	// DefaultIntValue is a default value for undefined env, configs etc
	DefaultIntValue = 0
	// DefaultInt64Value is a default value for undefined env, configs etc
	DefaultInt64Value = 0
	// DefaultUIntValue is a default value for undefined env, configs etc
	DefaultUIntValue = 0
	// DefaultUInt64Value is a default value for undefined env, configs etc
	DefaultUInt64Value = 0
)

// EventCallback is a callback function with data
type EventCallback func(interface{}) error

// Callback is a callback function
type Callback func() error

// Injector inject data/dependencies to object
type Injector interface {
	InjectTo(obj interface{}) error
}

// DataScope provide data provider
type DataScope interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Keys() ([]string, error)
}

// EventScope provide event interface
type EventScope interface {
	Trigger(int, interface{}) error
	On(int, EventCallback)
}

// Scope is global scope interface
type Scope interface {
	DataScope
	EventScope
	InjectTo(obj interface{}) error
}

// Module represent a app module
type Module interface {
	RegisterDependencies(App) error
	InitDependencies(App) error
	Run() error
}

// Bootstrap represent bootstrap of a app
type Bootstrap interface {
	Register(Module) error
	Init() error
	Run() error
}

// App represent a app
type App interface {
	Name() string
	Version() string
	RootFilespace() filesystem.Filespace
	GlobalScope() Scope
	EngineScope() Scope
	ArgsScope() Scope
	FilespaceScope() Scope
	ConfigScope() Scope
	DependencyScope() Scope
	AppScope() Scope
	CommandScope() Scope
	DependencyProvider() dependency.Provider
}
