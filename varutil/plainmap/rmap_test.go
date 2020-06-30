package plainmap

import (
	"testing"

	"github.com/goatcms/goatcore/varutil"
)

func TestRMapToPMapFromObject(t *testing.T) {
	t.Parallel()
	sourcemap := map[string]interface{}{
		"key1":   1,
		"keystr": "str",
		"keymap": map[string]interface{}{
			"key11": 11,
		},
	}

	outmap, err := RecursiveMapToPlainMap(sourcemap)
	if err != nil {
		t.Error(err)
		return
	}

	if outmap["key1"] != 1 {
		t.Error("key1 != 1")
	}
	if outmap["keystr"] != "str" {
		t.Error("keystr != str")
	}
	if outmap["keymap.key11"] != 11 {
		t.Error("keymap.key11 != 11")
	}
}

func TestRMapToPMapWithJson(t *testing.T) {
	t.Parallel()
	var sourcemap map[string]interface{}
	err := varutil.ObjectFromJSON(&sourcemap, "{\"key1\": \"1\", \"keymap\": {\"key11\": \"11\"}}")
	if err != nil {
		t.Error(err)
		return
	}

	outmap, err := RecursiveMapToPlainMap(sourcemap)
	if err != nil {
		t.Error(err)
		return
	}

	if outmap["key1"].(string) != "1" {
		t.Error("key1 != 1")
	}
	if outmap["keymap.key11"].(string) != "11" {
		t.Error("keymap.key11 != 11")
	}
}

func TestToRecursiveMap(t *testing.T) {
	t.Parallel()
	var (
		err      error
		result   map[string]interface{}
		node     map[string]interface{}
		ok       bool
		value    interface{}
		intvalue int
		svalue   string
	)
	if result, err = ToRecursiveMap(map[string]interface{}{
		"dep.key":  "a",
		"simpekey": "b",
		"intvalue": int(11),
	}); err != nil {
		t.Error(err)
		return
	}
	// check simpekey
	if value, ok = result["intvalue"]; !ok {
		t.Errorf("expected intvalue")
		return
	}
	if intvalue, ok = value.(int); !ok {
		t.Errorf("expected intvalue as int type")
		return
	}
	if intvalue != 11 {
		t.Errorf("expected intvalue equals to 11 and take %v", intvalue)
		return
	}
	// check simpekey
	if value, ok = result["simpekey"]; !ok {
		t.Errorf("expected simpekey")
		return
	}
	if svalue, ok = value.(string); !ok {
		t.Errorf("expected simpekey as string type")
		return
	}
	if svalue != "b" {
		t.Errorf("expected dep.key equals to 'b'")
		return
	}
	// check dep.key
	if value, ok = result["dep"]; !ok {
		t.Errorf("expected dep node")
		return
	}
	if node, ok = value.(map[string]interface{}); !ok {
		t.Errorf("expected dep as map[string]interface{} type")
		return
	}
	if value, ok = node["key"]; !ok {
		t.Errorf("expected dep.key value")
		return
	}
	if svalue, ok = value.(string); !ok {
		t.Errorf("expected dep.key as string value")
		return
	}
	if svalue != "a" {
		t.Errorf("expected dep.key equals to 'a'")
		return
	}

}

func TestStringMapToRecursiveMap(t *testing.T) {
	t.Parallel()
	var (
		err    error
		result map[string]interface{}
		node   map[string]interface{}
		ok     bool
		value  interface{}
		svalue string
	)
	if result, err = StringMapToRecursiveMap(map[string]string{
		"dep.key":  "a",
		"simpekey": "b",
	}); err != nil {
		t.Error(err)
		return
	}
	// check simpekey
	if value, ok = result["simpekey"]; !ok {
		t.Errorf("expected simpekey")
		return
	}
	if svalue, ok = value.(string); !ok {
		t.Errorf("expected simpekey as string type")
		return
	}
	if svalue != "b" {
		t.Errorf("expected dep.key equals to 'b'")
		return
	}
	// check dep.key
	if value, ok = result["dep"]; !ok {
		t.Errorf("expected dep node")
		return
	}
	if node, ok = value.(map[string]interface{}); !ok {
		t.Errorf("expected dep as map[string]interface{} type")
		return
	}
	if value, ok = node["key"]; !ok {
		t.Errorf("expected dep.key value")
		return
	}
	if svalue, ok = value.(string); !ok {
		t.Errorf("expected dep.key as string value")
		return
	}
	if svalue != "a" {
		t.Errorf("expected dep.key equals to 'a'")
		return
	}
}
