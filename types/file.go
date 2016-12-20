package types

import "github.com/goatcms/goat-core/filesystem"

// File is default file interface (use for many filesources)
type File interface {
	Filespace() filesystem.Filespace
	Path() string
}
