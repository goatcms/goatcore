package terminal

import (
	"fmt"
	"sync"

	"github.com/goatcms/goatcore/app"
)

// Arguments is structure implements app.Arguments
type Arguments struct {
	mu   sync.RWMutex
	data map[string]app.TerminalArgument
}

// NewArguments create new instance of app.Arguments
func NewArguments(args ...app.TerminalArgument) app.TerminalArguments {
	arguments := newArguments()
	for _, arg := range args {
		arguments.SetArgument(arg)
	}
	return arguments
}
func newArguments() *Arguments {
	return &Arguments{
		data: make(map[string]app.TerminalArgument),
	}
}

// Arguments return argument's names
func (arguments *Arguments) ArgumentNames() (keys []string) {
	arguments.mu.RLock()
	keys = make([]string, len(arguments.data))
	i := 0
	for key := range arguments.data {
		keys[i] = key
		i++
	}
	arguments.mu.RUnlock()
	return
}

// Argument return an argument by name
func (arguments *Arguments) Argument(name string) (argument app.TerminalArgument) {
	arguments.mu.RLock()
	argument = arguments.data[name]
	arguments.mu.RUnlock()
	return
}

// SetArgument define new argument
func (arguments *Arguments) SetArgument(args ...app.TerminalArgument) {
	arguments.mu.Lock()
	defer arguments.mu.Unlock()
	for _, argument := range args {
		name := argument.Name()
		if _, ok := arguments.data[name]; ok {
			panic(fmt.Sprintf("Argument %s is defined twice", name))
		}
		arguments.data[name] = argument
	}
}
