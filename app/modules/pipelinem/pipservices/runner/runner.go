package runner

import (
	"fmt"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// RunnerDeps is deps for runner
type RunnerDeps struct {
	SandboxesManager pipservices.SandboxesManager `dependency:"PipSandboxesManager"`
	TasksUnit        pipservices.TasksUnit        `dependency:"PipTasksUnit"`
	SharedMutex      commservices.SharedMutex     `dependency:"CommonSharedMutex"`
}

// Runner is piplines repository
type Runner struct {
	deps RunnerDeps
}

// NewRunner create a Runner instance
func NewRunner(deps RunnerDeps) *Runner {
	return &Runner{
		deps: deps,
	}
}

// Factory create a Runner instance
func Factory(dp dependency.Provider) (ri interface{}, err error) {
	var deps RunnerDeps
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	return pipservices.Runner(NewRunner(deps)), nil
}

// Run pipeline
func (runner *Runner) Run(pip pipservices.Pip) (err error) {
	var (
		tasksManager pipservices.TasksManager
		sandbox      pipservices.Sandbox
		task         pipservices.TaskWriter
	)
	if sandbox, err = runner.deps.SandboxesManager.Get(pip.Sandbox); err != nil {
		return err
	}
	if tasksManager, err = runner.deps.TasksUnit.FromScope(pip.Context.Scope); err != nil {
		return err
	}
	if task, err = tasksManager.Create(pip); err != nil {
		return err
	}
	go runner.runGo(tasksManager, sandbox, task)
	return nil
}

// Run pipeline
func (runner *Runner) runGo(tasksManager pipservices.TasksManager, sandbox pipservices.Sandbox, task pipservices.TaskWriter) {
	var (
		unlockHandler commservices.UnlockHandler
		err           error
		ctxScp        = task.IOContext().Scope()
	)
	defer task.Close()
	if err = runner.waitForTasks(task, tasksManager); err != nil {
		ctxScp.AppendError(err)
		return
	}
	task.SetStatus(fmt.Sprintf("wait for resources"))
	unlockHandler = runner.deps.SharedMutex.Lock(task.LockMap())
	defer unlockHandler.Unlock()
	task.SetStatus(fmt.Sprintf("execute"))
	if err = sandbox.Run(task.IOContext()); err != nil {
		task.SetStatus("fail")
		return
	}
	if err = ctxScp.Wait(); err != nil {
		task.SetStatus(fmt.Sprintf("fail"))
	}
}

// waitForTasks wait for all related task
func (runner *Runner) waitForTasks(task pipservices.TaskWriter, tasksManager pipservices.TasksManager) (err error) {
	var relatedTask pipservices.Task
	for _, taskName := range task.WaitList() {
		task.SetStatus(fmt.Sprintf("wait for %s task", taskName))
		if relatedTask = tasksManager.Get(taskName); relatedTask == nil {
			return goaterr.Errorf("Unknow task %s", taskName)
		}
		if err = relatedTask.Wait(); err != nil {
			return err
		}
	}
	return nil
}
