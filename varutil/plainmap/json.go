package plainmap

import (
	"sort"
	"strings"

	"github.com/buger/jsonparser"
)

// JSONToPlainStringMap conavert json map to a plain map
func JSONToPlainStringMap(data []byte) (map[string]string, error) {
	out := map[string]string{}
	if err := jsonToPlainStringMap("", out, data); err != nil {
		return nil, err
	}
	return out, nil
}

func jsonToPlainStringMap(resultKey string, result map[string]string, data []byte) error {
	return jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		var newResultKey string
		if resultKey != "" {
			newResultKey = resultKey + "." + string(key)
		} else {
			newResultKey = string(key)
		}
		switch dataType {
		case jsonparser.Object:
			return jsonToPlainStringMap(newResultKey, result, value)
		case jsonparser.String:
			result[newResultKey] = string(value)
		case jsonparser.Number:
			result[newResultKey] = string(value)
		}
		return nil
	})
}

func StringPlainmapToJSON(plainmap map[string]string) (json string, err error) {
	// sort keys
	keys := make([]string, len(plainmap))
	i := 0
	for k := range plainmap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	// prepare json
	_, json, err = stringPlainmapToJSON("", 0, keys, plainmap)
	if err != nil {
		return "", err
	}
	return "{" + json + "}", nil
}

func stringPlainmapToJSON(prefix string, index int, keys []string, plainmap map[string]string) (i int, json string, err error) {
	for i = index; i < len(keys); {
		fullkey := keys[i]
		if !strings.HasPrefix(fullkey, prefix) {
			return i, json, nil
		}
		diff := fullkey[len(prefix):]
		doti := strings.Index(diff, ".")
		if doti != -1 {
			var prejson string
			i, prejson, err = stringPlainmapToJSON(fullkey[:len(prefix)+doti+1], i, keys, plainmap)
			if err != nil {
				return 0, "", err
			}
			if json != "" {
				json += ","
			}
			json += formatStringJSON(diff[:doti]) + ":{" + prejson + "}"
		} else {
			if json != "" {
				json += ","
			}
			json += formatStringJSON(diff) + ":" + formatStringJSON(plainmap[fullkey])
			i++
		}
	}
	return i, json, nil
}
