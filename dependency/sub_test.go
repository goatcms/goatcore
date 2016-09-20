package dependency

import "testing"

func TestCreateSubDP(t *testing.T) {
	dp := NewSubProvider(NewProvider())
	if dp == nil {
		t.Error("Expected dependency.Provider object, got nil")
	}
}

func TestSubServiceStory(t *testing.T) {
	dp := NewSubProvider(NewProvider())
	if err := dp.AddService(MyDepName, MyDepFactory); err != nil {
		t.Error("Add service fail", err)
		return
	}
	instance, err := dp.Get(MyDepName)
	if err != nil {
		t.Error(err)
		return
	}
	if instance == nil {
		t.Error("instance can not be empty")
		return
	}

	o1 := &SimpleObject{}
	var loadable Loadable
	loadable = o1
	loadable.Load(&dp)
	if o1.Instance == nil {
		t.Error("object dependency instance can not be empty after Load")
		return
	}

	o2 := &SimpleObject{}
	o2.Load(&dp)
	if o2.Instance == nil {
		t.Error("object2 dependency instance can not be empty after Load")
		return
	}

	//provided dep sould be the same
	o1.Instance.Set(1)
	o2.Instance.Set(2)
	if (o1.Instance.Get() != o2.Instance.Get()) || (o1.Instance.Get() != 2) {
		t.Error("object1 and object2 should have the same vaues")
	}
}

func TestSubCircleStory(t *testing.T) {
	dp := NewSubProvider(NewProvider())
	if err := dp.AddService(MyCircleDepName, MyCircleDepFactory); err != nil {
		t.Error("Add service fail", err)
		return
	}
	_, err := dp.Get(MyCircleDepName)
	if err == nil {
		t.Error("should return error when dependencies are circled:", err)
	}
}

func TestSubOverwriteDefaultServiceStory(t *testing.T) {
	dp := NewSubProvider(NewProvider())
	if err := dp.AddDefaultService(MyDepName, MyDepFactory); err != nil {
		t.Error("Add default service fail", err)
		return
	}
	if err := dp.AddDefaultService(MyDepName, MyDepFactory); err == nil {
		t.Error("Set default service twice should return error", err)
		return
	}
	if err := dp.AddService(MyDepName, MyDepFactory); err != nil {
		t.Error("Add default service fail", err)
		return
	}
	if err := dp.AddService(MyDepName, MyDepFactory); err == nil {
		t.Error("Set service twice should return error", err)
		return
	}
}

func TestSubGet(t *testing.T) {
	parent := NewProvider()
	dp := NewSubProvider(parent)
	if err := parent.AddDefaultService(MyDepName, MyDepFactory); err != nil {
		t.Error("Add default service to parent fail", err)
		return
	}
	instance, err := dp.Get(MyDepName)
	if err != nil {
		t.Error(err)
		return
	}
	if instance == nil {
		t.Error("instance can not be empty")
		return
	}
}
