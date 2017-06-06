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
