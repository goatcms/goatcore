package fsi18loader

import (
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/fsloop"
	"github.com/goatcms/goatcore/i18n"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/varutil/plainmap"
)

// Load read and set translations from directory
func Load(fs filesystem.Filespace, basePath string, i18 i18n.I18N, scope app.Scope) error {
	loop := fsloop.NewLoop(&fsloop.LoopData{
		Filespace: fs,
		FileFilter: func(fs filesystem.Filespace, subPath string) bool {
			return strings.HasSuffix(subPath, ".json")
		},
		OnFile: func(fs filesystem.Filespace, subPath string) error {
			data, err := fs.ReadFile(subPath)
			if err != nil {
				return err
			}
			tmap, err := plainmap.JSONToPlainStringMap(data)
			if err != nil {
				return err
			}
			i18.Set(tmap)
			return nil
		},
	}, scope)
	loop.Run(basePath)
	loop.Wait()
	if errs := loop.Errors(); len(errs) != 0 {
		return goaterr.NewErrors(errs)
	}
	return nil
}
