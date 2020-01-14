package namespaces

import (
	"testing"

	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/app/scope"
)

func TestUnitDefaultNamespaceStory(t *testing.T) {
	t.Parallel()
	var (
		err       error
		unit      = NewUnit()
		scp       = scope.NewScope(scope.Params{})
		namespace pipservices.Namespaces
	)
	if namespace, err = unit.FromScope(scp, DefaultNamespace); err != nil {
		t.Error(err)
		return
	}
	if namespace != DefaultNamespace {
		t.Errorf("expected default namespaces")
	}
}

func TestUnitNamespaceStory(t *testing.T) {
	t.Parallel()
	var (
		err       error
		unit      = NewUnit()
		scp       = scope.NewScope(scope.Params{})
		namespace pipservices.Namespaces
	)
	if err = unit.Define(scp, "task", "lock"); err != nil {
		t.Error(err)
		return
	}
	if namespace, err = unit.FromScope(scp, DefaultNamespace); err != nil {
		t.Error(err)
		return
	}
	if namespace.Task() != "task" {
		t.Errorf("incorrect task: take '%s' and expected 'task'", namespace.Task())
	}
	if namespace.Lock() != "lock" {
		t.Errorf("incorrect lock: take '%s' and expected 'lock'", namespace.Lock())
	}
}

func TestUnitNamespaceBindStory(t *testing.T) {
	t.Parallel()
	var (
		err       error
		unit      = NewUnit()
		parent    = scope.NewScope(scope.Params{})
		child     = scope.NewScope(scope.Params{})
		namespace pipservices.Namespaces
	)
	if err = unit.Define(parent, "task", "lock"); err != nil {
		t.Error(err)
		return
	}
	if err = unit.Bind(parent, child); err != nil {
		t.Error(err)
		return
	}
	if namespace, err = unit.FromScope(child, DefaultNamespace); err != nil {
		t.Error(err)
		return
	}
	if namespace.Task() != "task" {
		t.Errorf("incorrect task: take '%s' and expected 'task'", namespace.Task())
	}
	if namespace.Lock() != "lock" {
		t.Errorf("incorrect lock: take '%s' and expected 'lock'", namespace.Lock())
	}
}
