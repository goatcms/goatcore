package fsi18loader

import (
	"testing"

	"github.com/goatcms/goatcore/i18n/i18mem"
)

func BenchmarkDepth(b *testing.B) {
	i18 := i18mem.NewI18N()
	fs, err := createBenchmarkFilespace(20, 5, 5)
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err = Load(fs, "./", i18, nil); err != nil {
				b.Error(err)
				return
			}
		}
	})
}

func BenchmarkLinear(b *testing.B) {
	i18 := i18mem.NewI18N()
	fs, err := createBenchmarkFilespace(2000, 1, 1)
	if err != nil {
		b.Error(err)
		return
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err = Load(fs, "./", i18, nil); err != nil {
				b.Error(err)
				return
			}
		}
	})
}
