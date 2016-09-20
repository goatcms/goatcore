package json

import (
	"path/filepath"

	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
)

// ReadJSON read data from json file to object
func ReadJSON(fs filesystem.Filespace, src string, object interface{}) error {
	jsonb, err := fs.ReadFile(src)
	if err != nil {
		return err
	}
	return varutil.ObjectFromJSON(object, string(jsonb))
}

// WriteJSON write data from object to json file
func WriteJSON(fs filesystem.Filespace, path string, object interface{}) error {
	json, err := varutil.ObjectToJSON(object)
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
