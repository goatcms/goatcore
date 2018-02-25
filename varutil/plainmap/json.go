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

// PlainStringMapToJSON convert plain map (which all valueas are string) to JSON
func PlainStringMapToJSON(plainmap map[string]string) (json string, err error) {
	// sort keys
	keys := make([]string, len(plainmap))
	i := 0
	for k := range plainmap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	// prepare json
	_, json, err = plainStringMapToJSON("", 0, keys, plainmap)
	if err != nil {
		return "", err
	}
	return "{" + json + "}", nil
}

func plainStringMapToJSON(prefix string, index int, keys []string, plainmap map[string]string) (i int, json string, err error) {
	for i = index; i < len(keys); {
		fullkey := keys[i]
		if !strings.HasPrefix(fullkey, prefix) {
			return i, json, nil
		}
		diff := fullkey[len(prefix):]
		doti := strings.Index(diff, ".")
		if doti != -1 {
			var prejson string
			i, prejson, err = plainStringMapToJSON(fullkey[:len(prefix)+doti+1], i, keys, plainmap)
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

// PlainStringMapToFormattedJSON convert plain map (which all valueas are string) to formatted JSON
func PlainStringMapToFormattedJSON(plainmap map[string]string) (json string, err error) {
	// sort keys
	keys := make([]string, len(plainmap))
	i := 0
	for k := range plainmap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	// prepare json
	_, json, err = plainStringMapToFormattedJSON("", 0, keys, plainmap, "  ", "  ")
	if err != nil {
		return "", err
	}
	return "{" + json + "\n}", nil
}

func plainStringMapToFormattedJSON(prefix string, index int, keys []string, plainmap map[string]string, sep string, spaces string) (i int, json string, err error) {
	for i = index; i < len(keys); {
		fullkey := keys[i]
		if !strings.HasPrefix(fullkey, prefix) {
			return i, json, nil
		}
		diff := fullkey[len(prefix):]
		doti := strings.Index(diff, ".")
		if doti != -1 {
			var prejson string
			i, prejson, err = plainStringMapToFormattedJSON(fullkey[:len(prefix)+doti+1], i, keys, plainmap, sep, spaces+sep)
			if err != nil {
				return 0, "", err
			}
			if json != "" {
				json += ","
			}
			json += "\n" + spaces + formatStringJSON(diff[:doti]) + ": {" + prejson + "\n" + spaces + "}"
		} else {
			if json != "" {
				json += ","
			}
			json += "\n" + spaces + formatStringJSON(diff) + ": " + formatStringJSON(plainmap[fullkey])
			i++
		}
	}
	return i, json, nil
}
