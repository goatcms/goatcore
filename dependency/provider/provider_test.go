package provider

import (
	"testing"

	"github.com/goatcms/goat-core/dependency"
)

type TestInterface interface {
	Value() int
	Test() bool
}

type One struct{}

func (o One) Value() int {
	return 1
}

func (o One) Test() bool {
	return true
}

func NewOne() TestInterface {
	return One{}
}

func OneFactory(dp dependency.Provider) (dependency.Instance, error) {
	return NewOne(), nil
}

type Two struct {
	One TestInterface `inject:"one"`
}

func (t Two) Value() int {
	return 2
}

func (t Two) Test() bool {
	return t.One != nil && t.One.Test()
}

func NewTwo() TestInterface {
	return &Two{}
}

func TwoFactory(dp dependency.Provider) (dependency.Instance, error) {
	two := NewTwo()
	if err := dp.InjectTo(two); err != nil {
		return nil, err
	}
	return two, nil
}

type Circle struct {
	Circle TestInterface `inject:"circle"`
}

type ObjectWithRequired struct {
	Some TestInterface `inject:"dep"`
}

type ObjectWithUnrequired struct {
	Some TestInterface `inject:"?dep"`
}

func NewCircle() dependency.Instance {
	return &Circle{}
}

func CircleFactory(dp dependency.Provider) (dependency.Instance, error) {
	ins := NewCircle()
	if err := dp.InjectTo(ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func TestDefaultFactory(t *testing.T) {
	dp := NewProvider(TagName)
	if err := dp.AddDefaultFactory("one", OneFactory); err != nil {
		t.Error("Add OneFactory fail", err)
		return
	}
	oneInstance, err := dp.Get("one")
	if err != nil {
		t.Error(err)
		return
	}
	if oneInstance == nil {
		t.Error("instance musn't be nil")
		return
	}
	one := oneInstance.(TestInterface)
	if one.Value() != 1 {
		t.Error("instance value should return 1")
		return
	}
}

func TestFactory(t *testing.T) {
	dp := NewProvider(TagName)
	if err := dp.AddDefaultFactory("one", TwoFactory); err != nil {
		t.Error("Add OneFactory fail", err)
		return
	}
	if err := dp.AddFactory("one", OneFactory); err != nil {
		t.Error("Add OneFactory fail", err)
		return
	}
	oneInstance, err := dp.Get("one")
	if err != nil {
		t.Error(err)
		return
	}
	if oneInstance == nil {
		t.Error("instance musn't be nil")
		return
	}
	one := oneInstance.(TestInterface)
	if one.Value() != 1 {
		t.Error("instance value should return 1")
		return
	}
}

func TestInjectTo(t *testing.T) {
	dp := NewProvider(TagName)
	if err := dp.AddDefaultFactory("one", OneFactory); err != nil {
		t.Error("Add OneFactory fail", err)
		return
	}
	if err := dp.AddDefaultFactory("two", TwoFactory); err != nil {
		t.Error("Add TwoFactory fail", err)
		return
	}
	twoInstance, err := dp.Get("two")
	if err != nil {
		t.Error(err)
		return
	}
	if twoInstance == nil {
		t.Error("instance musn't be nil")
		return
	}
	two := twoInstance.(TestInterface)
	if two.Value() != 2 {
		t.Error("instance should return 2")
		return
	}
	if !two.Test() {
		t.Error("instance.Test() fail - one didn't be injected")
		return
	}
}

func TestPreventCircle(t *testing.T) {
	dp := NewProvider(TagName)
	if err := dp.AddDefaultFactory("circle", CircleFactory); err != nil {
		t.Error(err)
		return
	}
	if _, err := dp.Get("circle"); err == nil {
		t.Error("should return error when dependencies are circled")
	}
}

func TestRequireInjection(t *testing.T) {
	dp := NewProvider(TagName)
	if err := dp.AddDefaultFactory("dep", OneFactory); err != nil {
		t.Error(err)
		return
	}
	o := &ObjectWithRequired{
		Some: nil,
	}
	if err := dp.InjectTo(o); err != nil {
		t.Error(err)
		return
	}
	if o.Some == nil {
		t.Errorf("o.Some can not be null after dependency injection")
	}
}

func TestUnrequireInjection(t *testing.T) {
	dp := NewProvider(TagName)
	if err := dp.AddDefaultFactory("dep", OneFactory); err != nil {
		t.Error(err)
		return
	}
	o := &ObjectWithUnrequired{
		Some: nil,
	}
	if err := dp.InjectTo(o); err != nil {
		t.Error(err)
		return
	}
	if o.Some == nil {
		t.Errorf("o.Some can not be null after dependency injection")
	}
}

func TestUnrequiredSkipInjection(t *testing.T) {
	dp := NewProvider(TagName)
	o := &ObjectWithUnrequired{
		Some: nil,
	}
	if err := dp.InjectTo(o); err != nil {
		t.Error(err)
		return
	}
	if o.Some != nil {
		t.Errorf("o.Some must be null for unrequired and undefined dependency")
	}
}
