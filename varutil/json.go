package varutil

import (
	"bytes"
	"encoding/json"
	"strings"
)

// ObjectFromJSON decode json data to object
func ObjectFromJSON(object interface{}, data string) error {
	jsonParser := json.NewDecoder(strings.NewReader(data))
	return jsonParser.Decode(object)
}

// DataFromJSON decode json data to object
/*func DataFromJSON(data []byte) (interface{}, error) {
	var outdata interface{}
	if err := json.Unmarshal(data, &outdata); err != nil {
		return nil, err
	}
	return outdata, nil
}*/

// ObjectToJSON convert object to json
func ObjectToJSON(object interface{}) (string, error) {
	bytes, err := JSONMarshal(object, true)
	return string(bytes), err
}

// JSONMarshal convert object to json (and prepare json to be more user friendly)
func JSONMarshal(v interface{}, unescape bool) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if unescape {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}
