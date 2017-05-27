package ormtypes

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goatcms/goatcore/varutil"
)

type LongIDGenerator struct {
	instance int64
}

func NewLongIDGenerator(instance int32) (*LongIDGenerator, error) {
	if instance&(^varutil.ThreeByteMask) != 0 {
		return nil, fmt.Errorf("instance id value is too big %v", instance)
	}
	return &LongIDGenerator{
		instance: int64(instance) & varutil.ThreeByteMask,
	}, nil
}

func (g LongIDGenerator) ID() LongID {
	return LongID{
		CreateAt: time.Until(time.Now()).Nanoseconds(),
		Hash:     (g.instance << 40) | (rand.Int63() & varutil.FiveByteMask),
	}
}
