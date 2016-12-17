package command

import (
	"fmt"
	"strings"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/app/goatapp"
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
func NewModule() *Module {
	return &Module{}
}

// RegisterDependency is init callback to register module dependencies
func (m *Module) RegisterDependency(goatapp.GoatApp) error {
	return nil
}

// InitDependency is init callback to inject dependencies inside module
func (m *Module) InitDependency(injector app.Injector) error {
	injector.InjectTo(m.dependencies)
	return nil
}

// Run is run event callback
func (m *Module) Run() error {
	deps := m.dependencies

	fmt.Println(deps.Name, " ", deps.Version)
	fmt.Println(deps.Welcome)
	fmt.Println()
	fmt.Println("Commands:")
	for _, key := range deps.CommandScope.Keys() {
		helpStr := deps.CommandScope.Get(key)
		if strings.HasPrefix(key, "help.") {
			fmt.Println(" ", key[5:], ": ", helpStr)
		}
	}
	if deps.Company != "" {
		fmt.Println(deps.Company)
	}
	fmt.Println("Developed with GoatCore (", deps.GoatVersion, ")")
	return nil
}
