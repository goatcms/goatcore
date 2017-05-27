package ormtypes_test

import (
	"testing"

	"github.com/goatcms/goatcore/db/ormtypes"
)

func TestLongIDGenerator(t *testing.T) {
	instanceID := int32(14567)
	generator, err := ormtypes.NewLongIDGenerator(instanceID)
	if err != nil {
		t.Error(err)
		return
	}
	id := generator.ID()
	if id.InstanceID() != instanceID {
		t.Errorf("randm must be 3 (and it is %v)", id.InstanceID())
	}
}
