package diskfs

import "github.com/goatcms/goat-core/dependency"

// BuildFilespaceFactory return local filespace builder for specific path
func BuildFilespaceFactory(path string) dependency.Factory {
	return func(dp dependency.Provider) (dependency.Instance, error) {
		return NewFilespace(path)
	}
}
