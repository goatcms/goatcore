package diskfs

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/filesystem"
)

func BuildFilespaceFactory(path string) dependency.Factory {
	return func(dp dependency.Provider) (dependency.Instance, error) {
		return NewFilespace(path)
	}
}
