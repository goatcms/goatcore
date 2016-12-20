package plainmap

// Any represent any type
type Any interface{}

// ToPlainMap conavert map to plain map
func ToPlainMap(source map[string]interface{}) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	if err := toPlainMapNode(out, source, "", ""); err != nil {
		return nil, err
	}
	return out, nil
}

func toPlainMapNode(out, source map[string]interface{}, basekey, separator string) error {
	for key, value := range source {
		outkey := basekey + separator + key
		switch v := value.(type) {
		case map[string]interface{}:
			if err := toPlainMapNode(out, v, basekey+separator+key, "."); err != nil {
				return err
			}
		default:
			out[outkey] = value
		}
	}
	return nil
}
