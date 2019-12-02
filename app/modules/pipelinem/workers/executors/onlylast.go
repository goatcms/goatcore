package executors

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/goatcms/webslots/slotsapp/common/config"
	"github.com/goatcms/webslots/slotsapp/services"
)

// OnlyLastExecutor is a task executor, which break execution when new signal come. Only last signal will executed
type OnlyLastExecutor struct {
	counter   int
	counterMU sync.Mutex
	executeMU sync.Mutex
	script    config.Script
	taskName  string
}

// NewOnlyLastExecutor create a simultaneously task executor
func NewOnlyLastExecutor(taskName string, script config.Script) services.ScriptExecutor {
	return &OnlyLastExecutor{
		counter:  0,
		script:   script,
		taskName: taskName,
	}
}

func (executor *OnlyLastExecutor) getCounter() int {
	executor.counterMU.Lock()
	defer executor.counterMU.Unlock()
	executor.counter++
	return executor.counter
}

// Run script
func (executor *OnlyLastExecutor) Run(response services.ResponseContext) (err error) {
	counter := executor.getCounter()
	return executor.execute(response, counter)
}

// Run script
func (executor *OnlyLastExecutor) execute(response services.ResponseContext, counter int) (err error) {
	executor.executeMU.Lock()
	defer executor.executeMU.Unlock()
	for _, command := range executor.script.Commands {
		var (
			buf bytes.Buffer
		)
		if len(response.Errors()) != 0 {
			return fmt.Errorf("Task break by context error (probably another task in cantext was failed)")
		}
		if executor.counter != counter {
			mgs := "OnlyLastExecutor: Break by a new signal"
			response.Print(executor.taskName, mgs)
			return fmt.Errorf(mgs)
		}
		commandString := fmt.Sprintf("@: %v %v", command.Command, strings.Join(command.Args, " "))
		response.Print(executor.taskName, commandString)
		cwd := executor.script.CWD
		if command.CWD != "" {
			cwd = command.CWD
		}
		cmd := exec.Command(command.Command, command.Args...)
		cmd.Dir = cwd
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		if err = cmd.Run(); err != nil {
			response.AddError(executor.taskName, err)
			return fmt.Errorf("script %v run fail %v %v %v", executor.taskName, command.Command, command.Args, err)
		}
		response.Print(executor.taskName, buf.String())
	}
	return nil
}
