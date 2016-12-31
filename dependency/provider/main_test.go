package provider

import "github.com/goatcms/goat-core/dependency"

const (
	MyDepName       = "MyDep"
	MyCircleDepName = "MyCircleDep"
	TagName         = "inject"
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

func (d *MyDep) Set(v int) {
	d.value = v
}

type MyCircleDepInterface interface {
}

type MyCircleDep struct {
	instance MyCircleDepInterface
}

func MyDepFactory(dp dependency.Provider) (interface{}, error) {
	return &MyDep{}, nil
}

func MyCircleDepFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := dp.Get(MyCircleDepName)
	if err != nil {
		return nil, err
	}
	return &MyCircleDep{
		instance: instance,
	}, nil
}

type SimpleObject struct {
	Instance *MyDep `inject:"MyDep"`
}
