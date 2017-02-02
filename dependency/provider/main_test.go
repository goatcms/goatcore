package provider

import (
	"fmt"
	"reflect"

	"github.com/goatcms/goat-core/dependency"
)

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

type TestInjector struct{}

// InjectTo inject dependencies to object
func (injector *TestInjector) InjectTo(obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)
		if !valueField.IsValid() {
			return fmt.Errorf("goatcore/dependency/TestInjector.InjectTo: %s is invalid", structField.Name)
		}
		if !valueField.CanSet() {
			return fmt.Errorf("goatcore/dependency/TestInjector.InjectTo: Cannot set %s field value", structField.Name)
		}
		switch valueField.Interface().(type) {
		case int:
			valueField.SetInt(2016)
		case string:
			valueField.SetString("teststring")
		}
	}
	return nil
}
