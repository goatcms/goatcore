package app

import (
	"github.com/goatcms/goatcore/filesystem"
)

// Module represent a app module
type Module interface {
	InitDependencies(App) error
	RegisterDependencies(App) error
	Run(App) error
}

// Bootstrap represent bootstrap of a app
type Bootstrap interface {
	Init() error
	Register(Module) error
	Run() error
}

// Bootstrap represent bootstrap of a app
type Version interface {
	Major() int
	Minor() int
	Path() int
	Suffix() string
	String() string
}

// App represent a app
type App interface {
	Injector
	AppHealthCheckers

	Arguments() []string
	DependencyProvider() DependencyProvider
	Filespaces() AppFilespaces
	IOContext() IOContext
	Name() string
	Scopes() AppScopes
	Terminal() TerminalManager
	Version() Version
}

type AppScopes interface {
	App() Scope
	Arguments() DataScope
	Config() DataScope
	Filespace() DataScope
}

type AppFilespaces interface {
	CWD() filesystem.Filespace
	Home() filesystem.Filespace
	Root() filesystem.Filespace
	Tmp() filesystem.Filespace
}

// HealthCheckerCallback is function to check application health
type HealthCheckerCallback func(App, Scope) (msg string, err error)

// AppHealthCheckers check application helth
type AppHealthCheckers interface {
	// HealthCheckerNames return HealthChecker's names
	HealthCheckerNames() []string
	// HealthChecker return an HealthChecker by name
	HealthChecker(name string) HealthCheckerCallback
	// SetHealthChecker set new health hecker. Panic if healthecker name is duplicated
	SetHealthChecker(name string, cb HealthCheckerCallback) (err error)
}
