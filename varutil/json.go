package varutil

import (
	"encoding/json"
	"github.com/goatcms/goat-core/filesystem"
	"io/ioutil"
	"os"
	"path/filepath"
	"bytes"
)

func ReadJson(src string, object interface{}) error {
	var err error
	src, err = filepath.Abs(src)
	if err != nil {
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
	if err != nil {
		return err
	}
	b, err := JSONMarshal(object, true)
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

func JSONMarshal(v interface{}, unescape bool) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if unescape {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}
