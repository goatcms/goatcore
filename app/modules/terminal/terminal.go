package terminal

import (
	"strings"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules"
	"github.com/goatcms/goatcore/app/scope/argscope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// IOTerminal is user communication interface
type IOTerminal struct {
	deps struct {
		App    app.App    `dependency:"App"`
		Input  app.Input  `dependency:"InputService"`
		Output app.Output `dependency:"OutputService"`
	}
	loopModeMutex sync.Mutex
	mutex         sync.Mutex
	isLoopMode    bool
}

// IOTerminalFactory create new IOTerminal instance
func IOTerminalFactory(dp dependency.Provider) (in interface{}, err error) {
	instance := &IOTerminal{
		isLoopMode: false,
	}
	if err = dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return modules.Terminal(instance), nil
}

// RunLoop run terminal loop
func (terminal *IOTerminal) RunLoop() (err error) {
	var (
		line string
		args []string
	)
	terminal.loopModeMutex.Lock()
	if terminal.isLoopMode {
		return goaterr.Errorf("terminal.IOTerminal: terminal loop is run many times")
	}
	terminal.isLoopMode = true
	terminal.loopModeMutex.Unlock()
	for {
		terminal.deps.Output.Printf("\n>")
		if line, err = terminal.deps.Input.ReadLine(); err != nil {
			terminal.isLoopMode = false
			return err
		}
		if args, err = varutil.SplitArguments(line); err != nil {
			terminal.isLoopMode = false
			return err
		}
		if len(args) != 0 && strings.ToLower(args[0]) == "close" {
			terminal.isLoopMode = false
			return app.CloseApp(terminal.deps.App)
		}
		if err = terminal.RunCommand(args); err != nil {
			terminal.isLoopMode = false
			return err
		}
	}
}

// RunString execute single command
func (terminal *IOTerminal) RunString(s string) (err error) {
	var args []string
	if args, err = varutil.SplitArguments(s); err != nil {
		return err
	}
	return terminal.RunCommand(args)
}

// RunCommand execute single command
func (terminal *IOTerminal) RunCommand(args []string) (err error) {
	var (
		commandName  string
		cbIns        interface{}
		cb           app.CommandCallback
		commandScope = terminal.deps.App.CommandScope()
		ctxScope     app.Scope
	)
	terminal.mutex.Lock()
	defer terminal.mutex.Unlock()
	if len(args) != 0 {
		commandName = strings.ToLower(args[0])
	}
	if commandName == "" {
		return HelpComamnd(terminal.deps.App, nil)
	}
	if cbIns, err = commandScope.Get("command." + commandName); err != nil || cbIns == nil {
		return goaterr.Errorf("Error: unknown command %s", commandName)
	}
	if ctxScope, err = argscope.NewScope(args, "command"); err != nil {
		return err
	}
	cb = cbIns.(app.CommandCallback)
	return cb(terminal.deps.App, ctxScope)
}
