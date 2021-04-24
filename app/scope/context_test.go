package scope

import (
	"context"
	"testing"
)

func TestContextScopeStopStory(t *testing.T) {
	t.Parallel()
	scp := NewContextScope()
	if scp.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	scp.Stop()
	if scp.IsDone() == false {
		t.Errorf("expeted doned scope")
	}
	if scp.Err() != nil {
		t.Errorf("expected scp.Err() = nil")
	}
	if len(scp.Errors()) != 0 {
		t.Errorf("expected scp.Errors() = nil")
	}
}

func TestContextScopeKillStory(t *testing.T) {
	t.Parallel()
	scp := NewContextScope()
	if scp.IsDone() == true {
		t.Errorf("expeted undone scope")
	}
	scp.Kill()
	if scp.IsDone() == false {
		t.Errorf("expeted doned scope")
	}
	if scp.Err() != context.Canceled {
		t.Errorf("expected scp.Err() = context.Canceled")
	}
	if len(scp.Errors()) == 0 {
		t.Errorf("expected scp.Errors() = [context.Canceled]")
	}
}

func TestContextWaitStory(t *testing.T) {
	t.Parallel()
	scp := NewContextScope()
	scp.AddTasks(1)
	go func() {
		scp.DoneTask()
	}()
	if err := scp.Wait(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
