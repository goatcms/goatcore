package namespaces

import (
	"testing"

	"github.com/goatcms/goatcore/app/scope"
)

func TestMutexStory(t *testing.T) {
	t.Parallel()
	var (
		err       error
		namespace = NewNamasepaces()
		scp       = scope.NewScope("test")
		value     string
	)
	if err = namespace.Set(scp, "pip", "project:goatcms"); err != nil {
		t.Error(err)
		return
	}
	if value, err = namespace.Get(scp, "pip"); err != nil {
		t.Error(err)
		return
	}
	if value != "project:goatcms" {
		t.Errorf("Expected 'project:goatcms' as project")
		return
	}
}

func TestMutexsetOnceStory(t *testing.T) {
	t.Parallel()
	var (
		err       error
		namespace = NewNamasepaces()
		scp       = scope.NewScope("test")
	)
	if err = namespace.Set(scp, "pip", "project:goatcms"); err != nil {
		t.Error(err)
		return
	}
	if err = namespace.Set(scp, "pip", "project:goatcore"); err == nil {
		t.Errorf("Expected error when overwrite namespace")
		return
	}
}
