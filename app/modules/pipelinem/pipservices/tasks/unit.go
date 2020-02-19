package tasks

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/dependency"
)

// UnitDeps contains dependencies required by Unit
type UnitDeps struct {
	NamespacesUnit pipservices.NamespacesUnit `dependency:"PipNamespacesUnit"`
}

// Unit connect scope with tasks
type Unit struct {
	deps UnitDeps
}

// NewUnit create a Unit instance
func NewUnit(deps UnitDeps) *Unit {
	return &Unit{
		deps: deps,
	}
}

// UnitFactory create a Unit instance
func UnitFactory(dp dependency.Provider) (ri interface{}, err error) {
	unit := &Unit{}
	if err = dp.InjectTo(&unit.deps); err != nil {
		return nil, err
	}
	return pipservices.TasksUnit(unit), nil
}

// FromScope return pipeline task manager from scope
func (unit *Unit) FromScope(scp app.Scope) (tasks pipservices.TasksManager, err error) {
	var (
		manager pipservices.TasksManager
		ins     interface{}
	)
	locker := scp.LockData()
	defer locker.Commit()
	if ins, err = locker.Get(scopeKey); err != nil {
		return nil, err
	}
	if ins != nil {
		return ins.(pipservices.TasksManager), nil
	}
	manager = NewTaskManager(unit.deps, scp)
	if err = locker.Set(scopeKey, manager); err != nil {
		return nil, err
	}
	return manager, nil
}

// BindScope bind scope to task manager
func (unit *Unit) BindScope(scp app.Scope, manager pipservices.TasksManager) (err error) {
	return scp.Set(scopeKey, manager)
}

// Clear remove pipelines scope data
func (unit *Unit) Clear(scp app.Scope) (err error) {
	return scp.Set(scopeKey, nil)
}
