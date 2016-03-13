package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
)

func IsScaffoldingDir(path string) bool {
	return filesystem.IsFile(path + scaffoldingPath)
}
