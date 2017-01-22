package terminal

import (
	"fmt"
	"strings"

	"github.com/goatcms/goat-core/app"
)

// ModuleDependencies is a dependencies for the module
type ModuleDependencies struct {
	Name         string    `app:"AppName"`
	Version      string    `app:"AppVersion"`
	Welcome      string    `app:"?AppWelcome"`
	Company      string    `app:"?AppCompany"`
	GoatVersion  string    `engine:"GoatVersion"`
	CommandScope app.Scope `global:"CommandScope"`
}

// Module is command unit
type Module struct {
	dependencies ModuleDependencies
}

// NewModule create new command module instance
func NewModule() app.Module {
	return app.Module(&Module{})
}

// RegisterDependency is init callback to register module dependencies
func (m *Module) RegisterDependencies(app.App) error {
	return nil
}

// InitDependency is init callback to inject dependencies inside module
func (m *Module) InitDependencies(app app.App) error {
	if err := app.GlobalScope().InjectTo(&m.dependencies); err != nil {
		return err
	}
	fmt.Println("(&m.dependencies", &m.dependencies)
	return nil
}

// Run is run event callback
func (m *Module) Run() error {
	deps := m.dependencies
	keys, err := deps.CommandScope.Keys()
	if err != nil {
		return err
	}

	fmt.Println(deps.Name, " ", deps.Version)
	if deps.Welcome != "" {
		fmt.Println(deps.Welcome)
	}
	fmt.Println()
	fmt.Println("Commands:")
	for _, key := range keys {
		if strings.HasPrefix(key, "help.") {
			helpStr, err := deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			fmt.Println(" ", key[5:], ": ", helpStr)
		}
	}
	if deps.Company != "" {
		fmt.Println(deps.Company)
	}

	fmt.Println()
	fmt.Println("Developed with GoatCore (", deps.GoatVersion, ")")
	return nil
}
