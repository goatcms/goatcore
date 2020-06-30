package plainmap

import (
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

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

// ToRecursiveMap conavert a plain map to a multi level tree
func ToRecursiveMap(source map[string]interface{}) (out map[string]interface{}, err error) {
	var (
		path []string
		node map[string]interface{}
	)
	out = make(map[string]interface{})
	for spath, value := range source {
		if spath == "" {
			return nil, goaterr.Errorf("ToRecursiveMap: empty key is no allowd")
		}
		path = strings.Split(spath, ".")
		if len(path) > 1 {
			if node, err = toRecursiveMapCreateNode(out, path[:len(path)-1]); err != nil {
				return nil, err
			}
		} else {
			node = out
		}
		node[path[len(path)-1]] = value
	}
	return out, nil
}

func toRecursiveMapCreateNode(rmap map[string]interface{}, path []string) (map[string]interface{}, error) {
	var (
		node = rmap
		v    interface{}
		ok   bool
	)
	if len(path) == 0 {
		return nil, goaterr.Errorf("ToRecursiveMap: empty path is not allowed")
	}
	for i := 0; i < len(path); i++ {
		key := path[i]
		if v, ok = node[key]; !ok {
			child := make(map[string]interface{})
			node[key] = child
			node = child
			continue
		}
		switch child := v.(type) {
		case map[string]interface{}:
			node = child
			continue
		default:
			return nil, goaterr.Errorf("%v is not a node", strings.Join(path[:i], "."))
		}
	}
	return node, nil
}

// StringMapToRecursiveMap conavert a plain string map to a multi level tree
func StringMapToRecursiveMap(source map[string]string) (out map[string]interface{}, err error) {
	var (
		path []string
		node map[string]interface{}
	)
	out = make(map[string]interface{})
	for spath, value := range source {
		if spath == "" {
			return nil, goaterr.Errorf("ToRecursiveMap: empty key is no allowd")
		}
		path = strings.Split(spath, ".")
		if len(path) > 1 {
			if node, err = toRecursiveMapCreateNode(out, path[:len(path)-1]); err != nil {
				return nil, err
			}
		} else {
			node = out
		}
		node[path[len(path)-1]] = value
	}
	return out, nil
}
