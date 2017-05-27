package ormtypes

import (
	"time"

	"github.com/goatcms/goatcore/varutil"
)

// meta
//a 8-byte value representing the nanoseconds since the Unix epoch,
//a 3-byte instance ID
//a 5-byte random

type LongID struct {
	CreateAt int64 `db:"idtime"`
	Hash     int64 `db:"idhash"`
}

func (id LongID) InstanceID() int32 {
	return int32(id.Hash >> 40)
}

func (id LongID) Random() int64 {
	return id.Hash & varutil.FiveByteMask
}

func (id LongID) Time() time.Time {
	return time.Unix(0, id.CreateAt)
}
