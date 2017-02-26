package terminal

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcore/app"
)

const (
	commandPrefix  = "help.command."
	argumentPrefix = "help.argument."
)

// Module is command unit
type Module struct {
	application  app.App
	dependencies struct {
		Name         string    `app:"AppName"`
		Version      string    `app:"AppVersion"`
		Welcome      string    `app:"?AppWelcome"`
		Company      string    `app:"?AppCompany"`
		GoatVersion  string    `engine:"GoatVersion"`
		CommandName  string    `argument:"?$1"`
		CommandScope app.Scope `dependency:"CommandScope"`
	}
}

// NewModule create new command module instance
func NewModule() app.Module {
	return &Module{}
}

// RegisterDependency is init callback to register module dependencies
func (m *Module) RegisterDependencies(a app.App) error {
	commandScope := a.CommandScope()
	commandScope.Set("command.h", m.Help)
	commandScope.Set("command.help", m.Help)
	return nil
}

// InitDependency is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) error {
	if err := a.DependencyProvider().InjectTo(&m.dependencies); err != nil {
		return err
	}
	m.application = a
	return nil
}

func (m *Module) Run() error {
	// header
	fmt.Println(m.dependencies.Name, " ", m.dependencies.Version)
	if m.dependencies.Company != "" {
		fmt.Println(m.dependencies.Company)
	}
	fmt.Printf("Supported by goatcore %s (%s) \n", m.dependencies.GoatVersion, "https://github.com/goatcms/goatcore")
	if m.dependencies.Welcome != "" {
		fmt.Printf("\n%s\n", m.dependencies.Welcome)
	}
	// content
	if m.dependencies.CommandName == "" {
		return m.Help(m.application)
	}
	commandIns, err := m.dependencies.CommandScope.Get("command." + m.dependencies.CommandName)
	if err != nil || commandIns == nil {
		fmt.Printf("Error: unknown command %s\n", m.dependencies.CommandName)
		return nil
	}
	command := commandIns.(func(app.App) error)
	return command(m.application)
}

func (m *Module) Help(app.App) error {
	keys, err := m.dependencies.CommandScope.Keys()
	if err != nil {
		return err
	}
	fmt.Printf("\nCommands:\n")
	for _, key := range keys {
		if strings.HasPrefix(key, commandPrefix) {
			helpStr, err := m.dependencies.CommandScope.Get(key)
			if err != nil {
				return err
			}
			fmt.Printf(" %s: %s\n", key[len(commandPrefix):], helpStr)
		}
	}
	fmt.Printf("\n\nArguments:\n")
	for _, key := range keys {
		if strings.HasPrefix(key, argumentPrefix) {
			helpStr, err := m.dependencies.CommandScope.Get(key)
			if err != nil {
				return err
			}
			fmt.Printf(" %s: %s\n", key[len(argumentPrefix):], helpStr)
		}
	}
	return nil
}
