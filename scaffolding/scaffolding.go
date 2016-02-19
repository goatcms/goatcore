package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
  "fmt"
  "os"
)

func Build(src, dest string, exclude []string) error {
  err := filesystem.ForAllFiles(src, func(file os.FileInfo, path string) {
    //TODO
  })
  if err != nil {
    return err
  }
  return nil
}
