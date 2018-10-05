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
	app  app.App
	deps struct {
		Name         string    `app:"AppName"`
		Version      string    `app:"AppVersion"`
		Welcome      string    `app:"?AppWelcome"`
		Company      string    `app:"?AppCompany"`
		GoatVersion  string    `engine:"GoatVersion"`
		CommandName  string    `argument:"?$1"`
		CommandScope app.Scope `dependency:"CommandScope"`

		Input  app.Input  `dependency:"InputService"`
		Output app.Output `dependency:"OutputService"`
	}
}

// NewModule create new command module instance
func NewModule() app.Module {
	return &Module{}
}

// RegisterDependencies is init callback to register module dependencies
func (m *Module) RegisterDependencies(a app.App) error {
	commandScope := a.CommandScope()
	commandScope.Set("command.h", m.Help)
	commandScope.Set("command.help", m.Help)
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func (m *Module) InitDependencies(a app.App) error {
	if err := a.DependencyProvider().InjectTo(&m.deps); err != nil {
		return err
	}
	m.app = a
	return nil
}

// Run start command line loop
func (m *Module) Run() error {
	// header
	m.deps.Output.Printf("%s %s\n", m.deps.Name, m.deps.Version)
	if m.deps.Company != "" {
		m.deps.Output.Printf("Develop by @%s all rights reserved\n", m.deps.Company)
	}
	m.deps.Output.Printf("Powered by GoatCore %s (%s)\n", m.deps.GoatVersion, "https://github.com/goatcms/goatcore")
	if m.deps.Welcome != "" {
		m.deps.Output.Printf("\n%s\n", m.deps.Welcome)
	}
	// content
	if m.deps.CommandName == "" {
		return m.Help(m.app)
	}
	commandIns, err := m.deps.CommandScope.Get("command." + m.deps.CommandName)
	if err != nil || commandIns == nil {
		fmt.Printf("Error: unknown command %s\n", m.deps.CommandName)
		return nil
	}
	command := commandIns.(func(app.App) error)
	return command(m.app)
}

// Help show help message
func (m *Module) Help(app.App) error {
	keys, err := m.deps.CommandScope.Keys()
	if err != nil {
		return err
	}
	isFirstCommand := true
	maxLength := 0
	for _, key := range keys {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}
	maxLength = maxLength - len(commandPrefix) + 1
	for _, key := range keys {
		if strings.HasPrefix(key, commandPrefix) {
			if isFirstCommand {
				m.deps.Output.Printf("\nCommands:\n")
				isFirstCommand = false
			}
			helpStr, err := m.deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			m.deps.Output.Printf("%s  %s\n", fixSpace(key[len(commandPrefix):], maxLength), helpStr)
		}
	}
	isFirstArgument := true
	for _, key := range keys {
		if strings.HasPrefix(key, argumentPrefix) {
			if isFirstArgument {
				m.deps.Output.Printf("\nArguments:\n")
				isFirstArgument = false
			}
			helpStr, err := m.deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			m.deps.Output.Printf("%11s  %s\n", fixSpace(key[len(argumentPrefix):], maxLength), helpStr)
		}
	}
	m.deps.Output.Printf("\n")
	return nil
}

func fixSpace(s string, l int) string {
	ll := l - len(s)
	prefix := make([]rune, ll)
	for i := range prefix {
		prefix[i] = ' '
	}
	return string(prefix) + s
}
