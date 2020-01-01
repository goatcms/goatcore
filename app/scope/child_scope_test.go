package scope

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
)

func TestParentScopeIsKillByChildWhenFail(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
	)
	t.Parallel()
	parentScope = NewScope("")
	childScope = NewChildScope(parentScope, parentScope, parentScope)
	defer childScope.Close()
	childScope.Kill()
	if parentScope.IsKilled() != true {
		t.Errorf("parent scope should be killed when chilld die")
	}
}

func TestChildScopeAppendError(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
	)
	t.Parallel()
	parentScope = NewScope("")
	childScope = NewChildScope(parentScope, parentScope, parentScope)
	defer childScope.Close()
	childScope.AppendError(fmt.Errorf("some err"))
	if len(childScope.Errors()) == 0 {
		t.Errorf("expected error in child scope and take: %v", childScope.Errors())
	}
	if len(parentScope.Errors()) == 0 {
		t.Errorf("expected error in parent scope and take: %v", parentScope.Errors())
	}
}

func TestChildScopeAppendErrors(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
	)
	t.Parallel()
	parentScope = NewScope("")
	childScope = NewChildScope(parentScope, parentScope, parentScope)
	defer childScope.Close()
	childScope.AppendErrors(fmt.Errorf("some err"), fmt.Errorf("some err2"))
	if len(childScope.Errors()) == 0 {
		t.Errorf("expected error in child scope and take: %v", childScope.Errors())
	}
	if len(parentScope.Errors()) == 0 {
		t.Errorf("expected error in parent scope and take: %v", parentScope.Errors())
	}
}

func TestChildScopeWait(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
		result      []int
		wg          = &sync.WaitGroup{}
	)
	t.Parallel()
	parentScope = NewScope("")
	childScope = NewChildScope(parentScope, parentScope, parentScope)
	defer childScope.Close()
	wg.Add(2)
	childScope.AddTasks(1)
	go (func() {
		defer childScope.DoneTask()
		defer wg.Done()
		result = append(result, 1)
	})()
	go (func() {
		defer wg.Done()
		childScope.Wait()
		result = append(result, 2)
	})()
	wg.Wait()
	if len(result) != 2 {
		t.Errorf("expected 2 result")
	}
	if result[0] != 1 || result[1] != 2 {
		t.Errorf("expected 1, 2 and take %v", result)
	}
}

func TestParentScopeWait(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
		err         error
	)
	t.Parallel()
	parentScope = NewScope("")
	childScope = NewChildScope(parentScope, parentScope, parentScope)
	go (func() {
		time.Sleep(1 * time.Millisecond)
		childScope.Close()
	})()
	if err = parentScope.Wait(); err != nil {
		t.Error(err)
		return
	}
}
