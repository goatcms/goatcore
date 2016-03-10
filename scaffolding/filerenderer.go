package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
	"io/ioutil"
	"os"
	"text/template"
)

type FileRenderer struct {
	Template *template.Template
	Data     interface{}
}

func NewFileRenderer(d Delimiters) (*FileRenderer, error) {
	r := &FileRenderer{}
	if err := r.Init(d); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *FileRenderer) Init(d Delimiters) error {
	r.Template = template.New("main")
	r.Template.Delims(d.Left, d.Right)
	return nil
}

func (r *FileRenderer) LoadTemplates(p string) error {
	loop := filesystem.DirLoop{
		OnFile: parseTemplateFactory(r.Template),
		OnDir:  nil,
		Filter: nil,
	}
	if err := loop.Run(p); err != nil {
		return err
	}
	return nil
}

func (r *FileRenderer) Render(src, dest string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	tmpl, err := r.Template.New(src).Parse(string(b))
	if err != nil {
		return err
	}
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()
	err = tmpl.Execute(file, r.Data)
	if err != nil {
		return err
	}
	return nil
}

func parseTemplateFactory(t *template.Template) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		tmpl := t.New(path)
		_, err = tmpl.Parse(string(b))
		if err != nil {
			return err
		}
		return nil
	}
}
