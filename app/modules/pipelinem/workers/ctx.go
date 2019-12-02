package workers

import (
	"fmt"
	"sync"

	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/webslots/slotsapp/common/config"
	"github.com/goatcms/webslots/slotsapp/services"
)

type taskContextRow struct {
	mu   sync.Mutex
	done bool
}

// TaskContext is a tasks executing system
type TaskContext struct {
	rowsMU    sync.Mutex
	rows      map[string]*taskContextRow
	executors map[string]services.ScriptExecutor
	tasks     map[string]*config.Task
	response  *ResponseContext
}

// NewTaskContext create a new context
func NewTaskContext(executors map[string]services.ScriptExecutor, tasks map[string]*config.Task, response *ResponseContext) *TaskContext {
	return &TaskContext{
		rows:      map[string]*taskContextRow{},
		executors: executors,
		tasks:     tasks,
		response:  response,
	}
}

// Response return response of the context
func (ctx *TaskContext) Response() *ResponseContext {
	return ctx.response
}

// Run task by name
func (ctx *TaskContext) row(taskName string) *taskContextRow {
	var (
		ok  bool
		row *taskContextRow
	)
	ctx.rowsMU.Lock()
	defer ctx.rowsMU.Unlock()
	if row, ok = ctx.rows[taskName]; !ok {
		row = &taskContextRow{
			done: false,
		}
		ctx.rows[taskName] = row
	}
	return row
}

// Run task by name
func (ctx *TaskContext) Run(taskName string) (err error) {
	if err = ctx.run(taskName, []string{}); err != nil {
		return err
	}
	return nil
}

// Run task by name
func (ctx *TaskContext) run(taskName string, callstack []string) (err error) {
	var (
		task           *config.Task
		scriptExecutor services.ScriptExecutor
		row            *taskContextRow
		ok             bool
		wg             = &sync.WaitGroup{}
	)
	if varutil.IsArrContainStr(callstack, taskName) {
		return fmt.Errorf("Executors: %s contains require cycle. Task %s is required many times: %v", callstack[0], taskName, callstack)
	}
	if task, ok = ctx.tasks[taskName]; !ok {
		return fmt.Errorf("Executors: %s task is unknown", taskName)
	}
	if scriptExecutor, ok = ctx.executors[taskName]; !ok {
		return fmt.Errorf("Executors: %s task executor is unknown", taskName)
	}
	row = ctx.row(taskName)
	// exit if task was executed
	row.mu.Lock()
	defer row.mu.Unlock()
	if row.done {
		return nil
	}
	// run extended tasks
	wg.Add(len(task.Extends))
	for _, extendedTask := range task.Extends {
		go func(extendedTask string) {
			defer wg.Done()
			if err = ctx.run(extendedTask, append(callstack, taskName)); err != nil {
				ctx.response.AddError(extendedTask, err)
				return
			}
		}(extendedTask)
	}
	wg.Wait()
	errors := ctx.response.Errors()
	if len(errors) != 0 {
		return goaterr.NewErrors(errors)
	}
	// run tak
	if err = scriptExecutor.Run(ctx.response); err != nil {
		return err
	}
	row.done = true
	return nil
}
