package tasks

import (
	"fmt"
	"sort"
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
	deps            UnitDeps
	tasksMU         sync.RWMutex
	tasks           map[string]*Task
	rootScope       app.Scope
	oBroadcast      app.BufferedBroadcast
	statusBroadcast app.BufferedBroadcast
	wg              sync.WaitGroup
}

// NewTaskManager create a Output instance
func NewTaskManager(deps UnitDeps, rootScope app.Scope) (manager *TaskManager) {
	manager = &TaskManager{
		deps:            deps,
		rootScope:       rootScope,
		tasks:           map[string]*Task{},
		oBroadcast:      bufferio.NewBroadcast(nil, nil),
		statusBroadcast: bufferio.NewBroadcast(nil, nil),
	}
	return manager
}

// Names return existed task names
func (manager *TaskManager) Names() (names []string) {
	manager.tasksMU.Lock()
	defer manager.tasksMU.Unlock()
	for key := range manager.tasks {
		names = append(names, key)
	}
	sort.Strings(names)
	return names
}

// Get return task by name
func (manager *TaskManager) Get(name string) (task pipservices.Task, ok bool) {
	manager.tasksMU.RLock()
	defer manager.tasksMU.RUnlock()
	task, ok = manager.tasks[name]
	return
}

// Create new task
func (manager *TaskManager) Create(pip pipservices.Pip) (result pipservices.TaskWriter, err error) {
	var (
		ok              bool
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
	childScope = scope.NewChild(parentScope, scope.ChildParams{
		Name: fmt.Sprintf("task:%s", taskname),
	})
	if err = manager.deps.NamespacesUnit.Define(childScope, childNamespaces); err != nil {
		childScope.Close()
		return nil, err
	}
	childScope.SetValue(scopeKey, manager)
	taskCtx = gio.NewIOContext(childScope, gio.NewIO(gio.IOParams{
		In:  pip.Context.In,
		Out: pip.Context.Out,
		Err: pip.Context.Err,
		CWD: pip.Context.CWD,
	}))
	task = NewTask(taskCtx, pip, manager.statusBroadcast, manager.doneTask)
	oLogger := gio.NewLogger(manager.oBroadcast, taskname)
	if err = task.OBroadcast().Add(oLogger); err != nil {
		childScope.Close()
		return nil, err
	}
	// add oLogger to oBroadcast
	manager.tasks[taskname] = task
	if err = manager.validWaitList([]string{taskname}, task, 100); err != nil {
		childScope.Close()
		return nil, err
	}
	if err = manager.rootScope.AddTasks(1); err != nil {
		childScope.Close()
		return nil, err
	}
	manager.wg.Add(1)
	return task, nil
}

func (manager *TaskManager) doneTask() {
	manager.wg.Done()
	manager.rootScope.DoneTask()
}

// OBroadcast return output broadcas
func (manager *TaskManager) OBroadcast() app.BufferedBroadcast {
	return manager.oBroadcast
}

// StatusBroadcast return output broadcas
func (manager *TaskManager) StatusBroadcast() app.BufferedBroadcast {
	return manager.statusBroadcast
}

// Summary return all tasks summary
func (manager *TaskManager) Summary(out app.Output) (err error) {
	var (
		task pipservices.Task
		ok   bool
	)
	names := manager.Names()
	sort.Strings(names)
	for _, name := range names {
		if task, ok = manager.Get(name); !ok {
			return goaterr.Errorf("Task %s undefines", name)
		}
		out.Printf("***************************\n")
		out.Printf("**   %s (%s)\n", name, task.Status())
		out.Printf("***************************\n")
		desc := task.Description()
		if desc != "" {
			out.Printf("\n'''%s'''\n\n", desc)
		}
		out.Printf(task.IOBroadcast().String())
		out.Printf("\n\n")
	}
	return nil
}

// Wait for all tasks
func (manager *TaskManager) Wait() (err error) {
	var errs []error
	manager.wg.Wait()
	manager.tasksMU.RLock()
	defer manager.tasksMU.RUnlock()
	for _, task := range manager.tasks {
		errs = goaterr.AppendError(errs, task.Wait())
	}
	return goaterr.ToError(errs)
}

func (manager *TaskManager) validWaitList(path []string, task pipservices.Task, counter int) (err error) {
	var childTask *Task
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
