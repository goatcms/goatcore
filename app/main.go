package app

import (
	"context"
	"io"
	"time"

	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/messages"
)

const (
	// KillEvent is kill action for current event
	KillEvent = iota
	// ErrorEvent is action for a error
	ErrorEvent = iota
	// CommitEvent is a action run when data is persist
	CommitEvent = iota
	// AfterCommitEvent is a action run after data persist
	AfterCommitEvent = iota
	// RollbackEvent is a action run when data is restore
	RollbackEvent = iota
	// AfterRollbackEvent is a action run when after restore
	AfterRollbackEvent = iota
	// BeforeCloseEvent is a action run before application/scope close
	BeforeCloseEvent = iota
	// CloseEvent is a action run to close application/scope
	CloseEvent = iota

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
	GoatVersionValue = "0.0.0dx"
	// GoatVersion is key to get goat version
	GoatVersion = "GoatVersion"

	// EngineTagName is a name for app vars / const injection
	EngineTagName = "engine"
	// ArgsTagName is a name for argument injection
	ArgsTagName = "argument"
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
	// GlobalScope is a scope for global events and data
	GlobalScope = "GlobalScope"

	// AppService is a default application service
	AppService = "App"
	// InputService is a default input service
	InputService = "InputService"
	// OutputService is a default output service
	OutputService = "OutputService"
	// ErrorService is a default error service
	ErrorService = "ErrorService"

	//RootFilespace is key for root filesystem.Filespace
	RootFilespace = "root"
	//TmpFilespace is key for tmp filespace
	TmpFilespace = "tmp"
	//CurrentFilespace is key for CWD (Current Working Directory) filespace
	CurrentFilespace = "current"

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

	// ENVArg is name default environment application argument
	ENVArg = "env"

	// DefaultEnv is name of default system environment
	DefaultEnv = "prod"
)

const (
	// DefaultDeadline is default dedline for application, scopes and lifecycles (it is 3 years)
	// It is maximum time we declared the application can work correctly
	DefaultDeadline = time.Hour * 24 * 365 * 3
)

// TypeConverter convert from type to other (string->int etc)
type TypeConverter func(interface{}) (interface{}, error)

// EventCallback is a callback function with data
type EventCallback func(interface{}) error

// Callback is a callback function
type Callback func() error

// CommandCallback is function call to run user command
type CommandCallback func(App, IOContext) (err error)

// HealthCheckerCallback is function to check application health
type HealthCheckerCallback func(App, Scope) (msg string, err error)

// Injector inject data/dependencies to object
type Injector interface {
	InjectTo(obj interface{}) error
}

// DataScope provide data provider
type DataScope interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Keys() ([]string, error)
	LockData() (transaction DataScopeLocker)
}

// DataScopeLocker provide data scope commitable interface
type DataScopeLocker interface {
	DataScope
	Commit() (err error)
}

// EventScope provide event interface
type EventScope interface {
	Trigger(int, interface{}) error
	On(int, EventCallback)
}

// ErrorScope provide error interface
type ErrorScope interface {
	Errors() []error
	ToError() error
	AppendError(err error)
	AppendErrors(err ...error)
}

// SyncScope provide sync interface
type SyncScope interface {
	ErrorScope

	Context() context.Context
	Kill()
	IsKilled() bool
	Wait() error
	AddTasks(delta int) (err error)
	DoneTask()
}

// Scope is global scope interface
type Scope interface {
	DataScope
	EventScope
	SyncScope
	dependency.Injector

	Close() (err error)
}

// Module represent a app module
type Module interface {
	RegisterDependencies(App) error
	InitDependencies(App) error
	Run(App) error
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
	Arguments() []string
	RootFilespace() filesystem.Filespace
	EngineScope() Scope
	ArgsScope() Scope
	FilespaceScope() Scope
	ConfigScope() Scope
	AppScope() Scope
	CommandScope() Scope
	DependencyProvider() dependency.Provider
	IOContext() IOContext
}

// Form represent a form data
type Form interface {
	Valid() (messages.MessageMap, error)
	Data() interface{}
}

// Input represent a standard input
type Input interface {
	io.Reader
	ReadWord() (string, error)
	ReadLine() (string, error)
}

// Output represent a standard output
type Output interface {
	io.Writer
	Printf(format string, a ...interface{}) error
}

// IO represent a standard input/output
type IO interface {
	In() Input
	Out() Output
	Err() Output
	CWD() filesystem.Filespace
}

// IOContext represent application execution context
type IOContext interface {
	IO() IO
	Scope() Scope
}
