package pipc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

// Clear run pip:clear command
func Clear(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			TasksUnit pipservices.TasksUnit `dependency:"PipTasksUnit"`
		}
	)
	if err = ctx.Scope().InjectTo(&deps); err != nil {
		return err
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	return deps.TasksUnit.Clear(ctx.Scope())
}
