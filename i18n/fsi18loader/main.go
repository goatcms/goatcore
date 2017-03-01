package fsi18loader

import (
	"strings"

	"github.com/buger/jsonparser"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/fsloop"
	"github.com/goatcms/goatcore/i18n"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

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
			tmap := map[string]string{}
			if err = LoadJSON("", tmap, data); err != nil {
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

func LoadJSON(resultKey string, result map[string]string, data []byte) error {
	return jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		var newResultKey string
		if resultKey != "" {
			newResultKey = resultKey + "." + string(key)
		} else {
			newResultKey = string(key)
		}
		switch dataType {
		case jsonparser.Object:
			return LoadJSON(newResultKey, result, value)
		case jsonparser.String:
			result[newResultKey] = string(value)
		}
		return nil
	})
}
