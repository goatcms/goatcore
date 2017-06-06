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
	templateFile1   = `{{define "list"}} {{join . ", "}}{{end}} `
	templateFile2   = `{{define "unusedDef1"}} {{join . ": "}}{{end}} `
	templateFile3   = `{{define "unusedDef2"}} {{join . "| "}}{{end}} `
	templateFile4   = `{{define "unusedDef3"}} {{join . "/ "}}{{end}} `
)

func TestLoadDefaultLayout(t *testing.T) {
	t.Parallel()
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	fs, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	// create test data
	if err := fs.MkdirAll("layouts/", 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("layouts/default/main.gohtml", []byte(masterTemplate), 0777); err != nil {
		t.Error(err)
		return
	}
	// test loop
	for ti := 0; ti < workers.AsyncTestReapeat; ti++ {
		provider := NewProvider(fs, goathtml.LayoutPath, goathtml.ViewPath, funcs)
		view, err := provider.Layout(goathtml.DefaultLayout, nil)
		if err != nil {
			t.Errorf("Errors: %v", err)
			return
		}
		buf := new(bytes.Buffer)
		if err := view.Execute(buf, guardians); err != nil {
			t.Error(err)
			return
		}
		result := buf.String()
		if !strings.Contains(result, "- Gamora") {
			t.Errorf("after render should contains '- Gamora' and other characters. It is: %v", result)
			return
		}
	}
}

func TestLoadViewWithDefaultLayout(t *testing.T) {
	t.Parallel()
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
	if err := fs.WriteFile("layouts/default/main.gohtml", []byte(masterTemplate), 0777); err != nil {
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
		view, err := provider.View(goathtml.DefaultLayout, "myview", nil)
		if err != nil {
			t.Errorf("Errors: %v", err)
			return
		}
		buf := new(bytes.Buffer)
		if err := view.Execute(buf, guardians); err != nil {
			t.Error(err)
			return
		}
		result := buf.String()
		if !strings.Contains(result, "Gamora,") {
			t.Errorf("layout template should be overwrited. result: %v", result)
			return
		}
	}
}

func TestLoadManyFiles(t *testing.T) {
	t.Parallel()
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
	if err := fs.WriteFile("layouts/default/main.gohtml", []byte(masterTemplate), 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("views/myview/file1.gohtml", []byte(templateFile1), 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("views/myview/file2.gohtml", []byte(templateFile2), 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("views/myview/file3.gohtml", []byte(templateFile3), 0777); err != nil {
		t.Error(err)
		return
	}
	if err := fs.WriteFile("views/myview/file4.gohtml", []byte(templateFile4), 0777); err != nil {
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
		result := buf.String()
		if !strings.Contains(result, "Gamora,") {
			t.Errorf("layout template should be overwrited. Result is: %v", result)
			return
		}
	}
}
