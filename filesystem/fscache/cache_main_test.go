package fscache

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
)

func TestCacheInterface(t *testing.T) {
	t.Parallel()
	_ = filesystem.Filespace(Cache{})
}
