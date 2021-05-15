package scope

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/injector"
	"github.com/goatcms/goatcore/app/scope/contextscope"
	"github.com/goatcms/goatcore/app/scope/datascope"
	"github.com/goatcms/goatcore/app/scope/eventscope"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/varutil/idutil"
)

// Params describe scope
type Params struct {
	ContextScope app.ContextScope
	DataScope    app.DataScope
	EventScope   app.EventScope
	Injector     app.Injector
	CID          string
	Name         string
}

// Scope is global scope interface
type Scope struct {
	app.ContextScope
	app.DataScope
	app.EventScope
	app.Injector

	cid        string
	closed     bool
	closeStack string
	mu         sync.Mutex
	parent     app.Scope
	sid        string
	wg         sync.WaitGroup
}

// New create new instance of scope
func New(params Params) app.Scope {
	var (
		sid string
	)
	if params.ContextScope == nil {
		params.ContextScope = contextscope.New()
	}
	if params.DataScope == nil {
		params.DataScope = datascope.New(make(map[interface{}]interface{}))
	}
	if params.EventScope == nil {
		params.EventScope = eventscope.New()
	}
	if params.Injector == nil {
		params.Injector = injector.NewNilInjector()
	}
	if params.Name != "" {
		sid = fmt.Sprintf("%s(%s)", idutil.StringID(), params.Name)
	} else {
		sid = idutil.StringID()
	}
	if params.CID == "" {
		params.CID = CorrlationID(params.Name)
	}
	return &Scope{
		ContextScope: params.ContextScope,
		DataScope:    params.DataScope,
		EventScope:   params.EventScope,
		Injector:     params.Injector,

		cid: params.CID,
		sid: sid,

		closed: false,
	}
}

// CID return correlation id (it is used to track multi machine tasks)
func (scp *Scope) CID() string {
	return scp.cid
}

// SID return scope id
func (scp *Scope) SID() string {
	return scp.sid
}

// Wait for finish of context execution
func (scp *Scope) Wait() error {
	scp.wg.Wait()
	return scp.Err()
}

// AddTasks add task to wait group
func (scp *Scope) AddTasks(delta int) (err error) {
	if scp.IsDone() {
		return ErrDoned
	}
	scp.wg.Add(delta)
	return nil
}

// DoneTask mark single task as done
func (scp *Scope) DoneTask() {
	scp.wg.Done()
}

// Close scope
func (scp *Scope) Close() (err error) {
	scp.preventDoubleClosed()
	scp.mu.Lock()
	defer scp.mu.Unlock()
	scp.preventDoubleClosed()
	scp.closeStack = string(debug.Stack())
	scp.closed = true
	scp.appendError(scp.EventScope.Trigger(app.BeforeCloseEvent, scp))
	if err = scp.Wait(); err != nil {
		scp.appendError(scp.EventScope.Trigger(app.BeforeRollbackEvent, scp))
		scp.appendError(scp.EventScope.Trigger(app.RollbackEvent, scp))
		scp.appendError(scp.EventScope.Trigger(app.AfterRollbackEvent, scp))
		scp.close()
		return scp.Err()
	}
	scp.appendError(scp.EventScope.Trigger(app.BeforeCommitEvent, scp))
	scp.appendError(scp.EventScope.Trigger(app.CommitEvent, scp))
	scp.appendError(scp.EventScope.Trigger(app.AfterCommitEvent, scp))
	scp.close()
	return scp.Err()
}

func (scp *Scope) close() {
	scp.appendError(scp.EventScope.Trigger(app.AfterCloseEvent, scp))
	if scp.parent != nil {
		scp.parent.DoneTask()
	}
	scp.DataScope = nil
	scp.EventScope = nil
	scp.Injector = nil
	scp.parent = nil
}

func (scp *Scope) preventDoubleClosed() {
	if scp.closed {
		panic(goaterr.Errorf("scope [%s] is closed at:\n%s\nAND AT:\n %s\n\n", scp.sid, scp.closeStack, string(debug.Stack())))
	}
}

// Kill scope
func (scp *Scope) Kill() {
	scp.preventClosed()
	scp.ContextScope.Kill()
	scp.appendError(scp.Trigger(app.KillEvent, nil))
}

// Stop scope
func (scp *Scope) Stop() {
	scp.preventClosed()
	scp.ContextScope.Stop()
	scp.appendError(scp.Trigger(app.StopEvent, nil))
}

// AppendError add an error to the scope
func (scp *Scope) AppendError(errs ...error) {
	scp.preventClosed()
	scp.appendError(errs...)
}

func (scp *Scope) appendError(errs ...error) {
	if len(errs) == 0 {
		return
	}
	i := 0
	filtred := make([]error, len(errs))
	for _, e := range errs {
		if e != nil {
			filtred[i] = e
			i++
		}
	}
	if i == 0 {
		return
	}
	filtred = filtred[:i]
	scp.ContextScope.AppendError(filtred...)
	scp.ContextScope.AppendError(scp.Trigger(app.ErrorEvent, filtred))
}

// Err return cumulative error if the scope context contains any error
func (scp *Scope) Err() error {
	errs := scp.ContextScope.Errors()
	if len(errs) == 0 {
		return nil
	}
	return goaterr.Wrap(fmt.Sprintf("%s scope errors:", scp.sid), errs...)
}

// GoContext return golang context object
func (scp *Scope) GoContext() context.Context {
	return scp
}

// BaseDataScope return unwrap DataScope object
// (help better utilize/recycle objects)
func (scp *Scope) BaseDataScope() app.DataScope {
	return scp.DataScope
}

// BaseEventScope return unwrap EventScope object
// (help better utilize/recycle objects)
func (scp *Scope) BaseEventScope() app.EventScope {
	return scp.EventScope
}

// BaseContextScope return unwrap ContextScope object
// (help better utilize/recycle objects)
func (scp *Scope) BaseContextScope() app.ContextScope {
	return scp.ContextScope
}

// BaseInjector return unwrap dependency injector object
// (help better utilize/recycle objects)
func (scp *Scope) BaseInjector() app.Injector {
	return scp.Injector
}

func (scp *Scope) preventClosed() {
	if scp.closed {
		panic(goaterr.Errorf("scope [%s] is closed at:\n%s", scp.sid, scp.closeStack))
	}
}
