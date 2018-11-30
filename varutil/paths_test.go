package varutil

import "testing"

func TestGOPath(t *testing.T) {
	t.Parallel()
	var result string
	result = GOPath("github.com/goatcms/goatcms/cmsapp/commands/servec")
	if result != "github.com/goatcms/goatcms" {
		t.Errorf("From github.com/goatcms/goatcms/cmsapp/commands/servec expected github.com/goatcms/goatcms path and take '%v'", result)
	}
	result = GOPath("https://github.com/goatcms/goatcms/cmsapp/commands/servec.git")
	if result != "github.com/goatcms/goatcms" {
		t.Errorf("From https://github.com/goatcms/goatcms/cmsapp/commands/servec.git expected github.com/goatcms/goatcms path and take '%v'", result)
	}
	result = GOPath("https://github.com/goatcms/goatcms.git")
	if result != "github.com/goatcms/goatcms" {
		t.Errorf("From https://github.com/goatcms/goatcms.git expected github.com/goatcms/goatcms path and take '%v'", result)
	}
	result = GOPath("some/wrong.path")
	if result != "" {
		t.Errorf("Should return empty string for incorrect path and take '%v'", result)
	}
}
