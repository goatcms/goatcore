package terminal

import (
	"fmt"
	"sync"

	"github.com/goatcms/goatcore/app"
)

// Commands is interface represent terminal data definition
type Commands struct {
	mu   sync.RWMutex
	data map[string]app.TerminalCommand
}

// NewCommands create new instance of app.Commands
func NewCommands(cmds ...app.TerminalCommand) app.TerminalCommands {
	commands := newCommands()
	for _, command := range cmds {
		commands.SetCommand(command)
	}
	return commands
}

func newCommands() (commands *Commands) {
	return &Commands{
		data: make(map[string]app.TerminalCommand),
	}
}

// Commands return command's names
func (commands *Commands) CommandNames() (keys []string) {
	commands.mu.RLock()
	keys = make([]string, len(commands.data))
	i := 0
	for key := range commands.data {
		keys[i] = key
		i++
	}
	commands.mu.RUnlock()
	return
}

// Command return a command by name
func (commands *Commands) Command(name string) (command app.TerminalCommand) {
	commands.mu.RLock()
	command = commands.data[name]
	commands.mu.RUnlock()
	return
}

// SetCommand set new comamnd
func (commands *Commands) SetCommand(cmds ...app.TerminalCommand) {
	commands.mu.Lock()
	defer commands.mu.Unlock()
	for _, command := range cmds {
		name := command.Name()
		if _, ok := commands.data[name]; ok {
			panic(fmt.Sprintf("Command %s is defined twice", name))
		}
		commands.data[name] = command
	}
	return
}
