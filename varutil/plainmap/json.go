package plainmap

import "github.com/buger/jsonparser"

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
