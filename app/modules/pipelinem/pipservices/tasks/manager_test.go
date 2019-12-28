package tasks

import (
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/gio/bufferio"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestUnitStory(t *testing.T) {
	t.Parallel()
	var (
		err     error
		mapp    app.App
		scp     = scope.NewScope("")
		manager pipservices.TasksManager
		task    pipservices.TaskWriter
	)
	if mapp, err = newApp(); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		TasksUnit pipservices.TasksUnit `dependency:"PipTasksUnit"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if manager, err = deps.TasksUnit.FromScope(scp); err != nil {
		t.Error(err)
		return
	}
	buffer := bufferio.NewBuffer()
	pipCtx := pipservices.PipContext{
		In:    gio.NewInput(strings.NewReader("")),
		Out:   bufferio.NewBufferOutput(buffer),
		Err:   bufferio.NewBufferOutput(buffer),
		Scope: scp,
	}
	if pipCtx.CWD, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if task, err = manager.Create(pipservices.Pip{
		Name:    "task",
		Context: pipCtx,
	}); err != nil {
		t.Error(err)
		return
	}
	if task == nil {
		t.Errorf("expected task and take nil")
	}
	// test logs
	if err = task.IOContext().IO().Out().Printf("sometext"); err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(manager.Logs(), "sometext") {
		t.Errorf("expected 'sometext' in logs")
	}
	if !strings.Contains(buffer.String(), "sometext") {
		t.Errorf("expected 'sometext' in task output")
	}
}
