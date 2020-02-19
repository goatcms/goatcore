package pipc

import (
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/varutil/varg"
)

// Run run pip:run command
func Run(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			Name    string `command:"?name"`
			Body    string `command:"?body"`
			RLock   string `command:"?rlock"`
			RWLock  string `command:"?wlock"`
			Wait    string `command:"?wait"`
			Sandbox string `command:"?sandbox"`
			Silent  string `command:"?silent" ,argument:"?silent"`

			Runner         pipservices.Runner         `dependency:"PipRunner"`
			NamespacesUnit pipservices.NamespacesUnit `dependency:"PipNamespacesUnit"`
		}
		out           app.Output
		erro          app.Output
		scpNamespaces pipservices.Namespaces
		lockMap       = commservices.LockMap{}
		wait          []string
		lockNamespace string
		silent        bool
	)
	if err = goaterr.ToError(goaterr.AppendError(nil,
		ctx.Scope().InjectTo(&deps),
		a.DependencyProvider().InjectTo(&deps),
		a.ArgsScope().InjectTo(&deps),
	)); err != nil {
		return err
	}
	deps.Name = strings.Trim(deps.Name, cutset)
	if deps.Name == "" {
		return goaterr.Errorf("pip:run Name is required")
	}
	if !namePattern.MatchString(deps.Name) {
		return goaterr.Errorf("pip:run Name '%s' is incorrect", deps.Name)
	}
	if silent, err = varg.MatchBool("silent argument", deps.Silent, true); err != nil {
		return err
	}
	deps.Body = strings.Trim(deps.Body, cutset)
	if deps.Body == "" {
		return goaterr.Errorf("pip:run Body is required")
	}
	if scpNamespaces, err = deps.NamespacesUnit.FromScope(ctx.Scope(), defaultNamespace); err != nil {
		return err
	}
	lockNamespace += scpNamespaces.Lock()
	if deps.RLock != "" {
		if err = markBoolMapForNamespace(deps.RLock, lockNamespace, commservices.LockR, lockMap); err != nil {
			return err
		}
	}
	if deps.RWLock != "" {
		if err = markBoolMapForNamespace(deps.RWLock, lockNamespace, commservices.LockRW, lockMap); err != nil {
			return err
		}
	}
	if deps.Wait != "" {
		waitPrefix := scpNamespaces.Task()
		if waitPrefix != "" {
			waitPrefix = waitPrefix + ":"
		}
		if wait, err = splitWaitNames(waitPrefix, deps.Wait); err != nil {
			return err
		}
	}
	ctxIO := ctx.IO()
	if silent {
		out = gio.NewNilOutput()
		erro = out
	} else {
		out = ctxIO.Out()
		erro = ctxIO.Err()
	}
	return deps.Runner.Run(pipservices.Pip{
		Context: pipservices.PipContext{
			In:    gio.NewInput(strings.NewReader(deps.Body)),
			Out:   out,
			Err:   erro,
			CWD:   ctxIO.CWD(),
			Scope: ctx.Scope(),
		},
		Name:       deps.Name,
		Namespaces: scpNamespaces,
		Sandbox:    deps.Sandbox,
		Lock:       lockMap,
		Wait:       wait,
	})
}
