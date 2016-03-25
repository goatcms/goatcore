package dependency_test

import (
	"github.com/goatcms/goat-core/dependency"
)

const (
	MyDepName = "MyDep"
)

type MyDepInterface interface {
	IsItOk() bool
	Set(int)
	Get() int
}

type MyDep struct {
	value int
}

func (d *MyDep) IsItOk() bool {
	return true
}

func (d *MyDep) Get() int {
	return d.value
}

func (d *MyDep) Set(v int)  {
	d.value = v
}

func MyDepFactory() (dependency.Instance, error) {
	return &MyDep{}, nil
}

type SimpleObject struct {
	Instance *MyDep
}

func (o *SimpleObject) Load(dp *dependency.Provider) error {
	ins, err := (*dp).Get(MyDepName)
	if err != nil {
		return err
	}
	o.Instance = ins.(*MyDep)
	return nil
}
