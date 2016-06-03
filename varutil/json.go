package varutil

import (
	"encoding/json"
	"bytes"
	"strings"
)

func ObjectFromJson(object interface{}, data string) error {
	jsonParser := json.NewDecoder(strings.NewReader(data))
	if err := jsonParser.Decode(object); err != nil {
		return err
	}
	return nil
}

func ObjectToJson(object interface{}) (string, error) {
	bytes, err := JSONMarshal(object, true)
	return string(bytes), err
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
