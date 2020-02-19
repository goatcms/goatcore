package scope

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app"
)

func TestParentScopeIsKillByChildParallerWhenFail(t *testing.T) {
	var (
		parentScope   app.Scope
		parallelScope app.Scope
	)
	t.Parallel()
	parentScope = NewScope(Params{})
	parallelScope = NewParallelScope(parentScope, Params{})
	defer parallelScope.Close()
	parallelScope.Kill()
	if parentScope.IsKilled() != true {
		t.Errorf("parent scope should be killed when chilld die")
	}
}

func TestParallelScopeAppendError(t *testing.T) {
	var (
		parentScope   app.Scope
		parallelScope app.Scope
	)
	t.Parallel()
	parentScope = NewScope(Params{})
	parallelScope = NewParallelScope(parentScope, Params{
		EventScope: parentScope,
		DataScope:  parentScope,
	})
	defer parallelScope.Close()
	parallelScope.AppendError(fmt.Errorf("some err"))
	if len(parallelScope.Errors()) == 0 {
		t.Errorf("expected error in child scope and take: %v", parallelScope.Errors())
	}
	if len(parentScope.Errors()) == 0 {
		t.Errorf("expected error in parent scope and take: %v", parentScope.Errors())
	}
}

func TestParallelScopeAppendErrors(t *testing.T) {
	var (
		parentScope   app.Scope
		parallelScope app.Scope
	)
	t.Parallel()
	parentScope = NewScope(Params{})
	parallelScope = NewParallelScope(parentScope, Params{
		EventScope: parentScope,
		DataScope:  parentScope,
	})
	defer parallelScope.Close()
	parallelScope.AppendErrors(fmt.Errorf("some err"), fmt.Errorf("some err2"))
	if len(parallelScope.Errors()) == 0 {
		t.Errorf("expected error in child scope and take: %v", parallelScope.Errors())
	}
	if len(parentScope.Errors()) == 0 {
		t.Errorf("expected error in parent scope and take: %v", parentScope.Errors())
	}
}

func TestParallelScopeWait(t *testing.T) {
	var (
		parentScope   app.Scope
		parallelScope app.Scope
		result        []int
		wg            = &sync.WaitGroup{}
	)
	t.Parallel()
	parentScope = NewScope(Params{})
	parallelScope = NewParallelScope(parentScope, Params{
		EventScope: parentScope,
		DataScope:  parentScope,
	})
	defer parallelScope.Close()
	wg.Add(2)
	parallelScope.AddTasks(1)
	go (func() {
		defer parallelScope.DoneTask()
		defer wg.Done()
		result = append(result, 1)
	})()
	go (func() {
		defer wg.Done()
		parallelScope.Wait()
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

func TestParentScopeNotWaitForParaller(t *testing.T) {
	var (
		parentScope   app.Scope
		parallelScope app.Scope
		err           error
	)
	t.Parallel()
	parentScope = NewScope(Params{})
	parallelScope = NewParallelScope(parentScope, Params{})
	result := ""
	parallelScope.AddTasks(1)
	go (func() {
		time.Sleep(10 * time.Millisecond)
		result += "2"
		parallelScope.DoneTask()
	})()
	result += "1"
	if err = parentScope.Wait(); err != nil {
		t.Error(err)
		return
	}
	if err = parallelScope.Wait(); err != nil {
		t.Error(err)
		return
	}
	if result != "12" {
		t.Errorf("expected 12 and take %v", result)
	}
}

func TestParallelScopeInjector(t *testing.T) {
	var (
		parentScope   app.Scope
		parallelScope app.Scope
		err           error
		result        struct {
			Value string `tagname:"key"`
		}
	)
	t.Parallel()
	parentScope = NewScope(Params{})
	ds := &DataScope{
		Data: map[string]interface{}{
			"key": "value",
		},
	}
	parallelScope = NewParallelScope(parentScope, Params{
		Injectors: []app.Injector{
			ds.Injector("tagname"),
		},
	})
	if err = parallelScope.InjectTo(&result); err != nil {
		t.Error(err)
		return
	}
	if result.Value != "value" {
		t.Errorf("expected result.Value equals to 'value'")
	}
}
