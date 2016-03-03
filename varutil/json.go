package varutil

import (
	"encoding/json"
	"github.com/goatcms/goat-core/filesystem"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadJson(src string, object interface{}) error {
	var err error
	src, err = filepath.Abs(src)
	if err!=nil {
		return err
	}
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(object); err != nil {
		return err
	}
	return nil
}

func WriteJson(path string, object interface{}) error {
	var err error
	path, err = filepath.Abs(path)
	if err!=nil {
		return err
	}
	b, err := json.Marshal(object)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if !filesystem.IsDir(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(path, b, 0777)
	if err != nil {
		return err
	}
	return nil
}
