package executors

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/goatcms/webslots/slotsapp/common/config"
	"github.com/goatcms/webslots/slotsapp/services"
)

// ParallelExecutor is a simultaneously task executor
type ParallelExecutor struct {
	script   config.Script
	taskName string
}

// NewParallelExecutor create a simultaneously task executor
func NewParallelExecutor(taskName string, script config.Script) services.ScriptExecutor {
	return &ParallelExecutor{
		script:   script,
		taskName: taskName,
	}
}

// Run script
func (executor *ParallelExecutor) Run(response services.ResponseContext) (err error) {
	for _, command := range executor.script.Commands {
		var (
			buf bytes.Buffer
		)
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
			response.AddError(executor.taskName, err)
			return fmt.Errorf("script %v run fail %v %v: %v", executor.taskName, command.Command, command.Args, err)
		}
		response.Print(executor.taskName, buf.String())
	}
	return nil
}
