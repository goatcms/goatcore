package plainmap

// RecursiveMapToPlainMap conavert multi level map to a plain map
func RecursiveMapToPlainMap(source map[string]interface{}) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	if err := recursiveMapToPlainMapNode(out, source, "", ""); err != nil {
		return nil, err
	}
	return out, nil
}

func recursiveMapToPlainMapNode(out, source map[string]interface{}, basekey, separator string) error {
	for key, value := range source {
		outkey := basekey + separator + key
		switch v := value.(type) {
		case map[string]interface{}:
			if err := recursiveMapToPlainMapNode(out, v, basekey+separator+key, "."); err != nil {
				return err
			}
		default:
			out[outkey] = value
		}
	}
	return nil
}
