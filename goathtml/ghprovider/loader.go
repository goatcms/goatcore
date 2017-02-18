package ghprovider

import (
	"html/template"

	"github.com/goatcms/goatcore/filesystem"
)

type TemplateLoader struct {
	template *template.Template
}

func NewTemplateLoader(template *template.Template) *TemplateLoader {
	return &TemplateLoader{
		template: template,
	}
}

func (loader *TemplateLoader) Load(fs filesystem.Filespace, subPath string) error {
	bytes, err := fs.ReadFile(subPath)
	if err != nil {
		return err
	}
	if _, err := loader.template.Parse(string(bytes)); err != nil {
		return err
	}
	return nil
}

func (loader *TemplateLoader) Template() *template.Template {
	return loader.template
}
