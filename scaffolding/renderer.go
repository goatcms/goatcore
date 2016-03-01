package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Renderer struct {
	BasePath string
	Data     RendererData
	Template *template.Template
}

type RendererData struct {
	Values  map[string]string `json:"values"`
	Secrets map[string]string `json:"secrets"`
	Root    interface{}       `json:"data"`
}

func NewRenderer(basePath string, delimiters Delimiters, data *RendererData) (*Renderer, error) {
	r := &Renderer{
		BasePath: basePath,
		Data:     *data,
	}

	if !strings.HasSuffix(basePath, "/") {
		basePath += "/"
	}

	r.Init(delimiters)
	return r, nil
}

func (r *Renderer) Init(d Delimiters) error {
	var err error
	r.Template = template.New("main")
	if err != nil {
		return err
	}
	r.Template.Delims(d.Left, d.Right)

	loop := filesystem.DirLoop{
		OnFile: parseTemplateFactory(r.Template),
		OnDir:  nil,
		Filter: nil,
	}
	if err = loop.Run(r.BasePath + TemplatesDir); err != nil {
		return err
	}

	return nil
}

func (r *Renderer) Render(src, dest string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	tmpl, err := r.Template.New(src).Parse(string(b))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, CreateMode)
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
