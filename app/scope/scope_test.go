package scope

import "testing"

func TestScopeWaitStory(t *testing.T) {
	t.Parallel()
	scp := New(Params{})
	scp.AddTasks(1)
	go func() {
		scp.DoneTask()
	}()
	if err := scp.Wait(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
