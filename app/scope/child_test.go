package scope

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope/datascope"
)

func TestParentScopeIsKillByChildWhenFail(t *testing.T) {
	var (
		parentScope = New(Params{})
		childScope  app.Scope
	)
	t.Parallel()
	childScope = NewChild(parentScope, ChildParams{})
	childScope.Kill()
	childScope.Close()
	if parentScope.IsDone() != true {
		t.Errorf("parent scope should be killed when chilld die")
	}
}

func TestChildScopeAppendError(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
	)
	t.Parallel()
	parentScope = New(Params{})
	childScope = NewChild(parentScope, ChildParams{
		EventScope: parentScope,
		DataScope:  parentScope,
	})
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
	parentScope = New(Params{})
	childScope = NewChild(parentScope, ChildParams{
		EventScope: parentScope,
		DataScope:  parentScope,
	})
	defer childScope.Close()
	childScope.AppendError(fmt.Errorf("some err"), fmt.Errorf("some err2"))
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
	parentScope = New(Params{})
	childScope = NewChild(parentScope, ChildParams{
		EventScope: parentScope,
		DataScope:  parentScope,
	})
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
	parentScope = New(Params{})
	childScope = NewChild(parentScope, ChildParams{})
	go (func() {
		time.Sleep(1 * time.Millisecond)
		childScope.Close()
	})()
	if err = parentScope.Wait(); err != nil {
		t.Error(err)
		return
	}
}

func TestChildScopeInjector(t *testing.T) {
	var (
		parentScope app.Scope
		childScope  app.Scope
		err         error
		result      struct {
			Value string `tagname:"key"`
		}
	)
	t.Parallel()
	parentScope = New(Params{})
	ds := datascope.New(map[interface{}]interface{}{
		"key": "value",
	})
	childScope = NewChild(parentScope, ChildParams{
		Injector: datascope.NewInjector("tagname", ds),
	})
	if err = childScope.InjectTo(&result); err != nil {
		t.Error(err)
		return
	}
	if result.Value != "value" {
		t.Errorf("expected result.Value equals to 'value'")
	}
}
