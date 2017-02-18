package ghprovider

import (
	"bytes"
	"html/template"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/workers"
)

const (
	masterTemplate  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	overlayTemplate = `{{define "list"}} {{join . ", "}}{{end}} `
)

func TestLoadViewWithDefaultLayout(t *testing.T) {
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
	}
	// create test data
	if err := fs.MkdirAll("layouts/", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.MkdirAll("views/", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("layouts/default/main.gohtml", []byte(overlayTemplate), 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("views/myview/main.gohtml", []byte(overlayTemplate), 0777); err != nil {
		t.Error(err)
		return
	}
	// test loop
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		provider := NewProvider(fs, goathtml.LayoutPath, goathtml.ViewPath, funcs)
		view, errs := provider.View(goathtml.DefaultLayout, "myview", nil)
		if errs != nil {
			t.Errorf("Errors: %v", errs)
			return
		}
		buf := new(bytes.Buffer)
		if err := view.Execute(buf, guardians); err != nil {
			t.Error(err)
			return
		}
		if strings.Contains(buf.String(), "Gamora,") {
			t.Errorf("layout template should be overwrited")
			return
		}
	}
}
