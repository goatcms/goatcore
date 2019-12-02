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

func TestSchedulerParallerRun(t *testing.T) {
	t.Parallel()
	var (
		response services.ResponseContext
		testFS   filesystem.Filespace
		err      error
		wg       *sync.WaitGroup
	)
	tasks := map[string]*config.Task{
		"parent": &config.Task{
			Name:     "parent",
			Executor: "fifo",
			Extends:  []string{"child1", "child2"},
			Locks:    []string{},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{},
		},
		"child1": &config.Task{
			Name:     "child1",
			Executor: "fifo",
			Extends:  []string{"base"},
			Locks:    []string{},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{},
		},
		"child2": &config.Task{
			Name:     "child2",
			Executor: "fifo",
			Extends:  []string{"base"},
			Locks:    []string{},
			Script: config.Script{
				Commands: []config.Command{},
			},
			Trigger: []string{},
		},
		"base": &config.Task{
			Name:     "base",
			Executor: "fifo",
			Extends:  []string{},
			Locks:    []string{},
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
	parentExecutor := NewMockedScriptExecutor()
	baseExecutor := NewMockedScriptExecutor()
	firstChildExecutor := NewMockedScriptExecutor()
	secondChildExecutor := NewMockedScriptExecutor()
	scheduler.scriptExecutors["base"] = baseExecutor
	scheduler.scriptExecutors["parent"] = parentExecutor
	scheduler.scriptExecutors["child1"] = firstChildExecutor
	scheduler.scriptExecutors["child2"] = secondChildExecutor
	if response, wg, err = scheduler.Run("someone", "parent"); err != nil {
		t.Error(err)
		return
	}
	wg.Wait()
	if len(response.Errors()) != 0 {
		t.Error(goaterr.NewErrors(response.Errors()))
		return
	}
	if !parentExecutor.Runned() {
		t.Errorf("parentExecutor should be runned")
	}
	if !baseExecutor.Runned() {
		t.Errorf("baseExecutor should be runned")
	}
	if !firstChildExecutor.Runned() {
		t.Errorf("firstChildExecutor should be runned")
	}
	if !secondChildExecutor.Runned() {
		t.Errorf("secondChildExecutor should be runned")
	}
	if !response.IsClosed() {
		t.Errorf("response should be closed")
	}
}
