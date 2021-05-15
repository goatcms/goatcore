package contextscope

import (
	"context"
	"testing"
)

func TestIsolatedContextScopeStopStory(t *testing.T) {
	t.Parallel()
	parent := New()
	isolated := NewIsolated(parent)
	if isolated.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	isolated.Stop()
	if parent.IsDone() == true {
		t.Errorf("expeted undone parent")
	}
	if isolated.IsDone() == false {
		t.Errorf("expeted done scope")
	}
	if isolated.Err() != nil {
		t.Errorf("expected scp.Err() = nil")
	}
	if len(isolated.Errors()) != 0 {
		t.Errorf("expected scp.Errors() = nil")
	}
}

func TestIsolatedContextScopeStopByParentStory(t *testing.T) {
	t.Parallel()
	parent := New()
	isolated := NewIsolated(parent)
	if isolated.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	parent.Stop()
	<-isolated.Done()
	if isolated.IsDone() == false {
		t.Errorf("expeted done scope")
	}
	if isolated.Err() != nil {
		t.Errorf("expected scp.Err() = nil")
	}
	if len(isolated.Errors()) != 0 {
		t.Errorf("expected scp.Errors() = nil")
	}
}

func TestParentScopeStopByIsolatedStory(t *testing.T) {
	t.Parallel()
	parent := New()
	isolated := NewIsolated(parent)
	if isolated.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	isolated.Stop()
	<-isolated.Done()
	if isolated.IsDone() == false {
		t.Errorf("expeted done scope")
	}
	if parent.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	if parent.Err() != nil {
		t.Errorf("expected scp.Err() = nil")
	}
	if len(parent.Errors()) != 0 {
		t.Errorf("expected scp.Errors() = nil")
	}
}

func TestIsolatedContextScopeKillStory(t *testing.T) {
	t.Parallel()
	parent := New()
	isolated := NewIsolated(parent)
	if isolated.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	isolated.Kill()
	<-isolated.Done()
	if parent.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	if isolated.IsDone() == false {
		t.Errorf("expeted done scope")
	}
	if isolated.Err() != context.Canceled {
		t.Errorf("expected scp.Err() = context.Canceled")
	}
	if len(isolated.Errors()) == 0 {
		t.Errorf("expected scp.Errors() = [context.Canceled]")
	}
}

func TestIsolatedContextScopeKillByParentStory(t *testing.T) {
	t.Parallel()
	parent := New()
	isolated := NewIsolated(parent)
	if isolated.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	parent.Kill()
	<-isolated.Done()
	if isolated.IsDone() == false {
		t.Errorf("expected done scope")
	}
	if isolated.Err() != context.Canceled {
		t.Errorf("expected scp.Err() = context.Canceled")
	}
	if len(isolated.Errors()) == 0 {
		t.Errorf("expected scp.Errors() = [context.Canceled]")
	}
}
