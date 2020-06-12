package pipc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Summary run pip:summary command
func Summary(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			TasksUnit pipservices.TasksUnit `dependency:"PipTasksUnit"`
		}
		taskManager pipservices.TasksManager
		task        pipservices.Task
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
		if task = taskManager.Get(taskName); task == nil {
			return goaterr.Errorf("Unknow task %s", taskName)
		}
		out.Printf("***************************\n")
		out.Printf("**   %s (%s)\n", taskName, task.Status())
		out.Printf("***************************\n")

		desc := task.Description()
		if desc != "" {
			out.Printf("\n'''%s'''\n\n", desc)
		}

		out.Printf(task.Logs())
		out.Printf("\n\n")
	}
	return nil
}
