package pipc

import (
	"strings"

	"github.com/goatcms/goatcore/app/injector"
	"github.com/goatcms/goatcore/app/scope"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/varutil/varg"
)

// Try run pip:try command
func Try(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			Name        string `command:"?name"`
			Description string `command:"?description"`
			TryBody     string `command:"?body"`
			SuccessBody string `command:"?success"`
			FailBody    string `command:"?fail"`
			FinallyBody string `command:"?finally"`
			Silent      string `command:"?silent" ,argument:"?silent"`

			Runner         pipservices.Runner         `dependency:"PipRunner"`
			NamespacesUnit pipservices.NamespacesUnit `dependency:"PipNamespacesUnit"`
			TasksUnit      pipservices.TasksUnit      `dependency:"PipTasksUnit"`
		}
		out           app.Output
		erro          app.Output
		scpNamespaces pipservices.Namespaces
		silent        bool
	)
	if err = goaterr.ToError(goaterr.AppendError(nil,
		ctx.Scope().InjectTo(&deps),
		a.InjectTo(&deps),
	)); err != nil {
		return err
	}
	deps.Name = strings.Trim(deps.Name, cutset)
	if deps.Name == "" {
		return goaterr.Errorf("pip:try Name is required")
	}
	if !namePattern.MatchString(deps.Name) {
		return goaterr.Errorf("pip:try Name '%s' is incorrect", deps.Name)
	}
	if silent, err = varg.MatchBool("silent argument", deps.Silent, true); err != nil {
		return err
	}
	deps.TryBody = strings.Trim(deps.TryBody, cutset)
	if deps.TryBody == "" {
		return goaterr.Errorf("pip:try Body is required")
	}
	deps.SuccessBody = strings.Trim(deps.SuccessBody, cutset)
	deps.FailBody = strings.Trim(deps.FailBody, cutset)
	deps.FinallyBody = strings.Trim(deps.FinallyBody, cutset)
	if scpNamespaces, err = deps.NamespacesUnit.FromScope(ctx.Scope(), defaultNamespace); err != nil {
		return err
	}
	scpNamespaces = namespaces.NewSubNamespaces(scpNamespaces, pipservices.NamasepacesParams{
		Task: deps.Name,
	})
	ctxIO := ctx.IO()
	if silent {
		out = gio.NewNilOutput()
		erro = out
	} else {
		out = ctxIO.Out()
		erro = ctxIO.Err()
	}
	// preload TaskManager to prevent bind it to separated scope (must be bind to root scope)
	parentScope := ctx.Scope()
	if _, err = deps.TasksUnit.FromScope(parentScope); err != nil {
		return err
	}
	if err = parentScope.AddTasks(1); err != nil {
		return err
	}
	// create independed scope for try body
	separatedScope := scope.New(scope.Params{
		DataScope:  parentScope,
		EventScope: parentScope,
		Injector:   injector.NewMultiInjector([]app.Injector{parentScope}),
	})
	if err = deps.Runner.Run(pipservices.Pip{
		Context: pipservices.PipContext{
			In:    gio.NewInput(strings.NewReader(deps.TryBody)),
			Out:   out,
			Err:   erro,
			CWD:   ctxIO.CWD(),
			Scope: separatedScope,
		},
		Name:        "body",
		Description: deps.Description,
		Namespaces:  scpNamespaces,
		Sandbox:     "self", // only self sandbox is supported
		Lock:        nil,    // lock is unsupported
		Wait:        nil,    // wait is unsupported
	}); err != nil {
		parentScope.DoneTask()
		return err
	}
	go func() {
		var catchErr error
		defer parentScope.DoneTask()
		catchErr = separatedScope.Wait()
		// run finally
		if deps.FinallyBody != "" {
			if err = deps.Runner.Run(pipservices.Pip{
				Context: pipservices.PipContext{
					In:    gio.NewInput(strings.NewReader(deps.FinallyBody)),
					Out:   out,
					Err:   erro,
					CWD:   ctxIO.CWD(),
					Scope: parentScope,
				},
				Name:        "finally",
				Description: "",
				Namespaces:  scpNamespaces,
				Sandbox:     "self", // only self sandbox is supported
				Lock:        nil,    // lock is unsupported
				Wait:        nil,    // wait is unsupported
			}); err != nil {
				parentScope.AppendError(err)
				return
			}
		}
		// run fail (if required)
		if deps.FailBody != "" && catchErr != nil {
			if err = deps.Runner.Run(pipservices.Pip{
				Context: pipservices.PipContext{
					In:    gio.NewInput(strings.NewReader(deps.FailBody)),
					Out:   out,
					Err:   erro,
					CWD:   ctxIO.CWD(),
					Scope: parentScope,
				},
				Name:        "fail",
				Description: catchErr.Error(),
				Namespaces:  scpNamespaces,
				Sandbox:     "self", // only self sandbox is supported
				Lock:        nil,    // lock is unsupported
				Wait:        nil,    // wait is unsupported
			}); err != nil {
				parentScope.AppendError(err)
				return
			}
		}
		// run success (if required)
		if deps.SuccessBody != "" && catchErr == nil {
			if err = deps.Runner.Run(pipservices.Pip{
				Context: pipservices.PipContext{
					In:    gio.NewInput(strings.NewReader(deps.SuccessBody)),
					Out:   out,
					Err:   erro,
					CWD:   ctxIO.CWD(),
					Scope: parentScope,
				},
				Name:        "success",
				Description: "",
				Namespaces:  scpNamespaces,
				Sandbox:     "self", // only self sandbox is supported
				Lock:        nil,    // lock is unsupported
				Wait:        nil,    // wait is unsupported
			}); err != nil {
				parentScope.AppendError(err)
				return
			}
		}
	}()
	return nil
}
