package app

import (
	"time"
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
	//HomeFilespace is key for home filespace
	HomeFilespace = "home"
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
