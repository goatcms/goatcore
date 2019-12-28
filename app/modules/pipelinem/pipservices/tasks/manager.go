package tasks

import (
	"strings"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/gio/bufferio"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// TaskManager control tasks
type TaskManager struct {
	deps       UnitDeps
	logs       *bufferio.Buffer
	logsOutput app.Output
	tasksMU    sync.RWMutex
	tasks      map[string]*Task
	scope      app.Scope
}

// NewTaskManager create a Output instance
func NewTaskManager(deps UnitDeps) (manager *TaskManager) {
	manager = &TaskManager{
		deps:  deps,
		logs:  bufferio.NewBuffer(),
		tasks: map[string]*Task{},
	}
	manager.logsOutput = bufferio.NewBufferOutput(manager.logs)
	return manager
}

// newTaskManager create a Output instance
func newTaskManager(deps UnitDeps) (manager pipservices.TasksManager) {
	return NewTaskManager(deps)
}

// Logs return logs
func (manager *TaskManager) Logs() string {
	return manager.logs.String()
}

// Names return existed task names
func (manager *TaskManager) Names() (names []string) {
	manager.tasksMU.Lock()
	defer manager.tasksMU.Unlock()
	for key := range manager.tasks {
		names = append(names, key)
	}
	return names
}

// Get return task by name
func (manager *TaskManager) Get(name string) (task pipservices.Task) {
	manager.tasksMU.RLock()
	defer manager.tasksMU.RUnlock()
	return manager.tasks[name]
}

// Create new task
func (manager *TaskManager) Create(pip pipservices.Pip) (result pipservices.TaskWriter, err error) {
	var (
		ok         bool
		repeatIO   app.IO
		childScope app.Scope
		taskCtx    app.IOContext
		task       *Task
	)
	if pip.Name == "" {
		return nil, goaterr.Errorf("Pip.Name is required")
	}
	if pip.Context.CWD == nil {
		return nil, goaterr.Errorf("Expected PipContext.CWD not nil")
	}
	if pip.Context.Err == nil {
		return nil, goaterr.Errorf("Expected PipContext.Err not nil")
	}
	if pip.Context.In == nil {
		return nil, goaterr.Errorf("Expected PipContext.In not nil")
	}
	if pip.Context.Out == nil {
		return nil, goaterr.Errorf("Expected PipContext.Out not nil")
	}
	if pip.Context.Scope == nil {
		return nil, goaterr.Errorf("Expected PipContext.Scope not nil")
	}
	manager.tasksMU.Lock()
	defer manager.tasksMU.Unlock()
	if _, ok = manager.tasks[pip.Name]; ok {
		return nil, goaterr.Errorf("Task '%s' is already defined", pip.Name)
	}
	outLogger := gio.NewLogger(manager.logsOutput, pip.Name)
	if repeatIO, err = bufferio.NewRepeatIO(pip.Context.In, gio.NewOutputBroadcast([]app.Output{
		outLogger,
		pip.Context.Out,
	}), gio.NewOutputBroadcast([]app.Output{
		outLogger,
		pip.Context.Err,
	}), pip.Context.CWD); err != nil {
		return nil, err
	}
	childScope = scope.NewScope("")
	if err = manager.deps.NamespacesUnit.Bind(pip.Context.Scope, childScope); err != nil {
		return nil, err
	}
	if taskCtx, err = gio.NewIOContext(childScope, repeatIO); err != nil {
		return nil, err
	}
	task = NewTask(taskCtx, pip)
	if err = manager.validWaitList([]string{pip.Name}, task, 100); err != nil {
		return nil, err
	}
	manager.tasks[pip.Name] = task
	return task, nil
}

func (manager *TaskManager) validWaitList(path []string, task pipservices.Task, counter int) (err error) {
	var childTask pipservices.Task
	if counter < 0 {
		return goaterr.Errorf("Too many depth. Your wait has to many depth %s", strings.Join(path, "->"))
	}
	for _, taskName := range task.WaitList() {
		if taskName == path[0] {
			return goaterr.Errorf("Detected waiting circle %s", strings.Join(path, "->"))
		}
		if childTask = manager.tasks[taskName]; childTask == nil {
			return goaterr.Errorf("Task '%s' undefined", taskName)
		}
		if err = manager.validWaitList(append(path, taskName), childTask, counter-1); err != nil {
			return err
		}
	}
	return nil
}
