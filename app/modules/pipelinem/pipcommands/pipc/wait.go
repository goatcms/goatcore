package pipc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Wait run pip:wait command
func Wait(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			TasksUnit pipservices.TasksUnit `dependency:"PipTasksUnit"`
		}
		taskManager pipservices.TasksManager
		task        pipservices.Task
		ok          bool
	)
	if err = ctx.Scope().InjectTo(&deps); err != nil {
		return err
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if taskManager, err = deps.TasksUnit.FromScope(ctx.Scope()); err != nil {
		return err
	}
	names := taskManager.Names()
	if len(names) == 0 {
		ctx.IO().Out().Printf("No task found")
		return nil
	}
	out := ctx.IO().Out()
	for _, taskName := range names {
		if task, ok = taskManager.Get(taskName); !ok {
			return goaterr.Errorf("Unknow task %s", taskName)
		}
		if !task.Done() {
			out.Printf("\n Wait for %s: %s ", taskName, task.Description())
			if err = task.Wait(); err != nil {
				return err
			}
		}
		out.Printf("\n [%s] %s... %s", taskName, task.Description(), task.Status())
		if len(task.Errors()) != 0 {
			for _, err = range task.Errors() {
				out.Printf("\n - %s", err.Error())
			}
		}
	}
	out.Printf("\n")
	return nil
}
