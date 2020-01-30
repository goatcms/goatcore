package tasks

import (
	"bytes"
	"strings"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/gio/bufferio"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
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
	rootScope  app.Scope
}

// NewTaskManager create a Output instance
func NewTaskManager(deps UnitDeps, rootScope app.Scope) (manager *TaskManager) {
	manager = &TaskManager{
		deps:      deps,
		rootScope: rootScope,
		logs:      bufferio.NewBuffer(),
		tasks:     map[string]*Task{},
	}
	manager.logsOutput = bufferio.NewBufferOutput(manager.logs)
	return manager
}

// newTaskManager create a Output instance
func newTaskManager(deps UnitDeps, rootScope app.Scope) (manager pipservices.TasksManager) {
	return NewTaskManager(deps, rootScope)
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
		ok              bool
		repeatIO        app.IO
		childScope      app.Scope
		taskCtx         app.IOContext
		task            *Task
		parentScope     = pip.Context.Scope
		childNamespaces pipservices.Namespaces
	)
	if pip.Name == "" {
		return nil, goaterr.Errorf("Pip.Name is required")
	}
	if pip.Namespaces == nil {
		return nil, goaterr.Errorf("Expected Pip.Namespaces not nil")
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
	if parentScope == nil {
		return nil, goaterr.Errorf("Expected PipContext.Scope not nil")
	}
	childNamespaces = namespaces.NewSubNamespaces(pip.Namespaces, pipservices.NamasepacesParams{
		Task: pip.Name,
	})
	taskname := childNamespaces.Task()
	manager.tasksMU.Lock()
	defer manager.tasksMU.Unlock()
	if _, ok = manager.tasks[taskname]; ok {
		return nil, goaterr.Errorf("Task '%s' is already defined", taskname)
	}
	if pip.LogsBuffer == nil {
		pip.LogsBuffer = &bytes.Buffer{}
	}
	if pip.Logs == nil {
		pip.Logs = gio.NewOutput(pip.LogsBuffer)
	}
	outLogger := gio.NewLogger(manager.logsOutput, taskname)
	repeatIO = bufferio.NewRepeatIO(gio.IOParams{
		In: pip.Context.In,
		Out: gio.NewOutputBroadcast([]app.Output{
			pip.Logs,
			outLogger,
			pip.Context.Out,
		}),
		Err: gio.NewOutputBroadcast([]app.Output{
			pip.Logs,
			outLogger,
			pip.Context.Err,
		}),
		CWD: pip.Context.CWD,
	})
	childScope = scope.NewChildScope(parentScope, scope.ChildParams{})
	if err = manager.deps.NamespacesUnit.Define(childScope, childNamespaces); err != nil {
		return nil, err
	}
	if err = childScope.Set(scopeKey, manager); err != nil {
		return nil, err
	}
	taskCtx = gio.NewIOContext(childScope, repeatIO)
	task = NewTask(taskCtx, pip, manager.rootScope.DoneTask)
	manager.tasks[taskname] = task
	if err = manager.validWaitList([]string{taskname}, task, 100); err != nil {
		return nil, err
	}
	manager.rootScope.AddTasks(1)
	return task, nil
}

func (manager *TaskManager) validWaitList(path []string, task pipservices.Task, counter int) (err error) {
	var childTask pipservices.Task
	if counter < 0 {
		return goaterr.Errorf("Too many depth. Your wait has to many depth %s", strings.Join(path, "->"))
	}
	if len(task.WaitList()) == 0 {
		return nil
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
