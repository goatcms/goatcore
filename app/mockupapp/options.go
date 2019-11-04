package mockupapp

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
)

// MockupOptions contains settings for MockupApp
type MockupOptions struct {
	Name    string
	Version string

	Args   []string
	Input  app.Input
	Output app.Output

	RootFilespace    filesystem.Filespace
	TMPFilespace     filesystem.Filespace
	CurrentFilespace filesystem.Filespace

	EngineScope     app.Scope
	ArgsScope       app.Scope
	FilespaceScope  app.Scope
	ConfigScope     app.Scope
	DependencyScope app.Scope
	AppScope        app.Scope
	CommandScope    app.Scope

	DP dependency.Provider
}
