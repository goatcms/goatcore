package app

import (
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
)

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
	HomeFilespace() filesystem.Filespace
	EngineScope() Scope
	ArgsScope() Scope
	FilespaceScope() Scope
	ConfigScope() Scope
	AppScope() Scope
	CommandScope() Scope
	DependencyProvider() dependency.Provider
	IOContext() IOContext
}
