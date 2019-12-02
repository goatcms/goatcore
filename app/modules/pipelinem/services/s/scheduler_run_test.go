package workers

import (
	"sync"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/webslots/slotsapp/common/config"
	"github.com/goatcms/webslots/slotsapp/services"
)

func TestSchedulerRun(t *testing.T) {
	t.Parallel()
	var (
		response services.ResponseContext
		testFS   filesystem.Filespace
		err      error
		wg       *sync.WaitGroup
	)
	tasks := map[string]*config.Task{
		"task1": &config.Task{
			Name:     "task1",
			Executor: "fifo",
			Extends:  []string{"task2"},
			Locks:    []string{"diskSpace1"},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{"triggerDeployTask1"},
		},
		"task2": &config.Task{
			Name:     "task2",
			Executor: "fifo",
			Extends:  []string{},
			Locks:    []string{"diskSpace2"},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{"triggerDeployTask2"},
		},
		"triggerDeployTask1": &config.Task{
			Name:     "triggerDeployTask1",
			Executor: "fifo",
			Extends:  []string{},
			Locks:    []string{"server1"},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{},
		},
		"triggerDeployTask2": &config.Task{
			Name:     "triggerDeployTask2",
			Executor: "fifo",
			Extends:  []string{},
			Locks:    []string{"server1"},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{},
		},
	}
	if testFS, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	resourcesManager := NewResourcesManager()
	scheduler, err := NewScheduler(tasks, resourcesManager, testFS)
	if err != nil {
		t.Error(err)
		return
	}
	task1Executor := NewMockedScriptExecutor()
	task2Executor := NewMockedScriptExecutor()
	deploy1Executor := NewMockedScriptExecutor()
	deploy2Executor := NewMockedScriptExecutor()
	scheduler.scriptExecutors["task1"] = task1Executor
	scheduler.scriptExecutors["task2"] = task2Executor
	scheduler.scriptExecutors["triggerDeployTask1"] = deploy1Executor
	scheduler.scriptExecutors["triggerDeployTask2"] = deploy2Executor
	if response, wg, err = scheduler.Run("someone", "task1"); err != nil {
		t.Error(err)
		return
	}
	wg.Wait()
	if len(response.Errors()) != 0 {
		t.Error(goaterr.NewErrors(response.Errors()))
		return
	}
	if !task1Executor.Runned() {
		t.Errorf("Task1 should be runned")
	}
	if !task2Executor.Runned() {
		t.Errorf("Task2 should be runned")
	}
	if !deploy1Executor.Runned() {
		t.Errorf("triggerDeployTask1 deploy task should be runned")
	}
	if !deploy2Executor.Runned() {
		t.Errorf("triggerDeployTask2 deploy task should be runned")
	}
	if !response.IsClosed() {
		t.Errorf("response should be closed")
	}
}
