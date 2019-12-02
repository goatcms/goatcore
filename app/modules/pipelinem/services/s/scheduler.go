package workers

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/webslots/slotsapp/common/config"
	"github.com/goatcms/webslots/slotsapp/services"
	"github.com/goatcms/webslots/slotsapp/services/workers/executors"
)

type schedulerRow struct {
	config   *config.Task
	locks    []string
	triggers []string
}

// Scheduler is a script executing system
type Scheduler struct {
	resourcesManager *ResourcesManager
	rows             map[string]*schedulerRow
	scriptExecutors  map[string]services.ScriptExecutor
	tasks            map[string]*config.Task
	fs               filesystem.Filespace
}

// Factory create a executing system instance
func Factory(dp dependency.Provider) (ri interface{}, err error) {
	var (
		deps struct {
			TasksStorage services.TasksStorage `dependency:"TasksStorage"`
			FS           filesystem.Filespace  `filespace:"root"`
		}
		fs filesystem.Filespace
	)
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	if err = deps.FS.MkdirAll("data/logs", filesystem.DefaultUnixDirMode); err != nil {
		return nil, err
	}
	if fs, err = deps.FS.Filespace("data/logs"); err != nil {
		return nil, err
	}
	return NewScheduler(deps.TasksStorage.GetAll(), NewResourcesManager(), fs)
}

// NewScheduler create a Scheduler instance
func NewScheduler(tasks map[string]*config.Task, resourcesManager *ResourcesManager, fs filesystem.Filespace) (ins *Scheduler, err error) {
	ins = &Scheduler{
		resourcesManager: resourcesManager,
		tasks:            tasks,
		rows:             map[string]*schedulerRow{},
		scriptExecutors:  map[string]services.ScriptExecutor{},
		fs:               fs,
	}
	if err = ins.load(); err != nil {
		return nil, err
	}
	return ins, nil
}

func (scheduler *Scheduler) load() (err error) {
	if err = scheduler.loadConfig(); err != nil {
		return err
	}
	scheduler.loadRowsResources()
	scheduler.loadRowsTriggers()
	return nil
}

func (scheduler *Scheduler) loadConfig() error {
	var (
		scriptExecutor services.ScriptExecutor
	)
	for _, task := range scheduler.tasks {
		executorName := strings.ToLower(task.Executor)
		switch executorName {
		case "fifo":
			scriptExecutor = executors.NewFIFOExecutor(task.Name, task.Script)
		case "parallal":
			scriptExecutor = executors.NewParallelExecutor(task.Name, task.Script)
		case "onlylast":
			scriptExecutor = executors.NewOnlyLastExecutor(task.Name, task.Script)
		}
		scheduler.scriptExecutors[task.Name] = scriptExecutor
		scheduler.rows[task.Name] = &schedulerRow{
			config: task,
		}
	}
	return nil
}

func (scheduler *Scheduler) loadRowsResources() {
	for _, row := range scheduler.rows {
		if row.locks == nil {
			scheduler.loadRowResources(row)
		}
	}
}

func (scheduler *Scheduler) loadRowResources(row *schedulerRow) (locks []string, err error) {
	var (
		ok                  bool
		relatedRow          *schedulerRow
		relatedRowResources []string
	)
	if row.locks != nil {
		return row.locks, nil
	}
	locks = row.config.Locks
	for _, taskName := range row.config.Extends {
		if relatedRow, ok = scheduler.rows[taskName]; !ok {
			return nil, fmt.Errorf("%s task is related to undefined task '%s'", row.config.Name, taskName)
		}
		if relatedRowResources, err = scheduler.loadRowResources(relatedRow); err != nil {
			return nil, err
		}
		for _, s := range relatedRowResources {
			if !varutil.IsArrContainStr(locks, s) {
				locks = append(locks, s)
			}
		}
	}
	sort.Strings(locks)
	row.locks = locks
	return locks, nil
}

func (scheduler *Scheduler) loadRowsTriggers() {
	for _, row := range scheduler.rows {
		if row.triggers == nil {
			scheduler.loadRowTriggers(row)
		}
	}
}

func (scheduler *Scheduler) loadRowTriggers(row *schedulerRow) (triggers []string, err error) {
	var (
		ok                 bool
		relatedRow         *schedulerRow
		relatedRowTriggers []string
	)
	if row.triggers != nil {
		return row.triggers, nil
	}
	triggers = row.config.Trigger
	for _, taskName := range row.config.Extends {
		if relatedRow, ok = scheduler.rows[taskName]; !ok {
			return nil, fmt.Errorf("%s task is related to undefined task '%s'", row.config.Name, taskName)
		}
		if relatedRowTriggers, err = scheduler.loadRowTriggers(relatedRow); err != nil {
			return nil, err
		}
		for _, s := range relatedRowTriggers {
			if !varutil.IsArrContainStr(triggers, s) {
				triggers = append(triggers, s)
			}
		}
	}
	sort.Strings(triggers)
	row.triggers = triggers
	return triggers, nil
}

// Run task by name
func (scheduler *Scheduler) Run(triggeredBy, taskName string) (response services.ResponseContext, waitGroup *sync.WaitGroup, err error) {
	waitGroup = &sync.WaitGroup{}
	waitGroup.Add(1)
	if response, err = scheduler.runWithWaitGroup(triggeredBy, taskName, waitGroup); err != nil {
		waitGroup.Done()
		return nil, nil, err
	}
	return response, waitGroup, nil
}

func (scheduler *Scheduler) runWithWaitGroup(triggeredBy, taskName string, waitGroup *sync.WaitGroup) (response services.ResponseContext, err error) {
	var (
		context         *TaskContext
		responseContext *ResponseContext
	)
	if responseContext, err = NewResponseContext(scheduler.fs, taskName, triggeredBy); err != nil {
		response.AddError(taskName, err)
		waitGroup.Done()
		return response, err
	}
	responseContext.WaitGroup().Add(1)
	context = NewTaskContext(scheduler.scriptExecutors, scheduler.tasks, responseContext)
	go scheduler.runInContext(triggeredBy, taskName, waitGroup, context)
	return responseContext, nil
}

func (scheduler *Scheduler) runInContext(triggeredBy, taskName string, waitGroup *sync.WaitGroup, context *TaskContext) {
	var (
		response = context.Response()
		row      *schedulerRow
		ok       bool
		err      error
	)
	defer waitGroup.Done()
	defer response.WaitGroup().Done()
	defer response.Close()
	if row, ok = scheduler.rows[taskName]; !ok {
		response.AddError(taskName, fmt.Errorf("%s task is undefined", taskName))
	}
	//response.WaitGroup().Add(1)
	if err = scheduler.lockAndRun(row, context); err != nil {
		response.AddError(taskName, err)
	}
	if len(response.Errors()) == 0 {
		triggerer := fmt.Sprintf("%v Task", taskName)
		waitGroup.Add(len(row.triggers))
		for _, triggerTask := range row.triggers {
			go scheduler.runWithWaitGroup(triggerer, triggerTask, waitGroup)
			response.Print("MAIN", fmt.Sprintf("Trigger %v task", triggerTask))
		}
	}
}

func (scheduler *Scheduler) lockAndRun(row *schedulerRow, context *TaskContext) (err error) {
	//defer response.WaitGroup().Done()
	scheduler.resourcesManager.LockAll(row.locks)
	defer scheduler.resourcesManager.UnlockAll(row.locks)
	return context.Run(row.config.Name)
}
