package simpletype

import "testing"

func TestBaseMetaType(t *testing.T) {
	ct := NewContentType(map[string]string{
		"maxlen": "10",
	})

	if ct.HasAttribute("maxlen") == false {
		t.Errorf("maxlen attribute should be defined")
	}

	if ct.GetAttribute("maxlen") != "10" {
		t.Errorf("maxlen attribute should be equels to '10', it is %v", ct.GetAttribute("maxlen"))
	}

}
