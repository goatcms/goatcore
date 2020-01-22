package pipc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

// Logs run pip:logs command
func Logs(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			TasksUnit pipservices.TasksUnit `dependency:"PipTasksUnit"`
		}
		taskManager pipservices.TasksManager
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
	out := ctx.IO().Out()
	out.Printf(taskManager.Logs())
	return nil
}
