package workers

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/webslots/slotsapp/common/config"
)

func TestSchedulerLoad(t *testing.T) {
	t.Parallel()
	var (
		testFS filesystem.Filespace
		err    error
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
	if scheduler.rows["task1"].config.Name != "task1" {
		t.Errorf("task1 Name should be equals to 'task1''")
		return
	}
	// test resources in task1
	if len(scheduler.rows["task1"].locks) != 2 {
		t.Errorf("task1 row locks should contains 2 eleemnts")
		return
	}
	if !varutil.IsArrContainStr(scheduler.rows["task1"].locks, "diskSpace1") {
		t.Errorf("task1 row locks should contains 'diskSpace1' resource")
		return
	}
	if !varutil.IsArrContainStr(scheduler.rows["task1"].locks, "diskSpace2") {
		t.Errorf("task1 row locks should contains 'diskSpace2' resource")
		return
	}
	// test locks in task2
	if len(scheduler.rows["task2"].locks) != 1 {
		t.Errorf("task2 row locks should contains 1 eleemnts")
		return
	}
	if !varutil.IsArrContainStr(scheduler.rows["task2"].locks, "diskSpace2") {
		t.Errorf("task1 row locks should contains 'diskSpace1' resource")
		return
	}
	// test triggers of task1
	if len(scheduler.rows["task1"].triggers) != 2 {
		t.Errorf("task1 row triggers should contains 2 eleemnts")
		return
	}
	if !varutil.IsArrContainStr(scheduler.rows["task1"].triggers, "triggerDeployTask1") {
		t.Errorf("task1 row triggers should contains 'triggerDeployTask1' resource")
		return
	}
	if !varutil.IsArrContainStr(scheduler.rows["task1"].triggers, "triggerDeployTask2") {
		t.Errorf("task1 row triggers should contains 'triggerDeployTask2' resource")
		return
	}
	// test triggers in task2
	if len(scheduler.rows["task2"].triggers) != 1 {
		t.Errorf("task2 row triggers should contains 1 eleemnts")
		return
	}
	if !varutil.IsArrContainStr(scheduler.rows["task2"].triggers, "triggerDeployTask2") {
		t.Errorf("task1 row locks should contains 'triggerDeployTask2' resource")
		return
	}
}
