package ormtypes_test

import (
	"testing"
	"time"

	"github.com/goatcms/goatcore/db/ormtypes"
)

func TestID(t *testing.T) {
	id := ormtypes.LongID{
		CreateAt: 1234,
		Hash:     (3 << 40) + 1,
	}
	if id.Random() != 1 {
		t.Errorf("randm must be 1 (and it is %v)", id.Random())
	}
	if id.InstanceID() != 3 {
		t.Errorf("randm must be 3 (and it is %v)", id.InstanceID())
	}
	expectedTime := time.Unix(0, 1234)
	if id.Time() != expectedTime {
		t.Errorf("incorect time")
	}
}
