package mockupapp

import (
	"io"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
)

// MockupOptions contains settings for MockupApp
type MockupOptions struct {
	Name    string
	Version string

	Input io.Reader

	Args []string

	RootFilespace filesystem.Filespace
	TMPFilespace  filesystem.Filespace
	HomeFilespace filesystem.Filespace

	EngineScope    app.Scope
	ArgsScope      app.Scope
	FilespaceScope app.Scope
	ConfigScope    app.Scope
	AppScope       app.Scope
	CommandScope   app.Scope

	DP dependency.Provider
}
