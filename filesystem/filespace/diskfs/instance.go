package diskfs

import "github.com/goatcms/goatcore/dependency"

// BuildFilespaceFactory return local filespace builder for specific path
func BuildFilespaceFactory(path string) dependency.Factory {
	return func(dp dependency.Provider) (interface{}, error) {
		return NewFilespace(path)
	}
}
