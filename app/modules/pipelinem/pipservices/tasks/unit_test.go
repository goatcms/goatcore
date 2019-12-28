package tasks

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

func TestManagertory(t *testing.T) {
	t.Parallel()
	var (
		err  error
		mapp app.App
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
}
