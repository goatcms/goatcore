package json

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
	"path/filepath"
)

func ReadJson(fs filesystem.Filespace, src string, object interface{}) error {
	jsonb, err := fs.ReadFile(src)
	if err != nil {
		return err
	}
	return varutil.ObjectFromJson(object, string(jsonb))
}

func WriteJson(fs filesystem.Filespace, path string, object interface{}) error {
	json, err := varutil.ObjectToJson(object)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if !fs.IsDir(dir) {
		if err := fs.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
	return fs.WriteFile(path, []byte(json), 0777)
}
