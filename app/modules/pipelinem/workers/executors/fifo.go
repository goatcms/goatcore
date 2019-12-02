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

// FIFOExecutor is a first in first out script executor
type FIFOExecutor struct {
	mu       sync.Mutex
	script   config.Script
	taskName string
}

// NewFIFOExecutor create new first in first out signal reducer
func NewFIFOExecutor(taskName string, script config.Script) services.ScriptExecutor {
	return &FIFOExecutor{
		script:   script,
		taskName: taskName,
	}
}

// Run script
func (executor *FIFOExecutor) Run(response services.ResponseContext) (err error) {
	executor.mu.Lock()
	defer executor.mu.Unlock()
	for _, command := range executor.script.Commands {
		var buf bytes.Buffer
		if len(response.Errors()) != 0 {
			return fmt.Errorf("Task break by context error (probably another task in cantext was failed)")
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
			return fmt.Errorf("script run fail %v %v: %v %v", command.Command, command.Args, err, buf.String())
		}
		response.Print(executor.taskName, buf.String())
	}
	return nil
}
