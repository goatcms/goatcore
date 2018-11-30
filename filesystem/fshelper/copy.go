package fshelper

import (
	"path"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/fsloop"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Copy clone files from sourcefs to destfs
func Copy(srcfs, destfs filesystem.Filespace, filterFN filesystem.LoopFilter) (err error) {
	loop := fsloop.NewLoop(&fsloop.LoopData{
		Filespace: srcfs,
		DirFilter: filterFN,
		OnDir: func(fs filesystem.Filespace, subPath string) error {
			return destfs.MkdirAll(subPath, filesystem.DefaultUnixDirMode)
		},
		OnFile: func(fs filesystem.Filespace, subPath string) (err error) {
			if err = destfs.MkdirAll(path.Dir(subPath), filesystem.DefaultUnixDirMode); err != nil {
				return err
			}
			if err = StreamCopy(srcfs, destfs, subPath); err != nil {
				return err
			}
			return nil
		},
		Consumers:  1,
		Producents: 1,
	}, nil)
	loop.Run("")
	loop.Wait()
	if len(loop.Errors()) != 0 {
		return goaterr.NewErrors(loop.Errors())
	}
	return err
}
