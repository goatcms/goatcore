package ghprovider

import (
	"fmt"
	"html/template"
	"sync"

	"github.com/goatcms/goatcore/filesystem"
)

type TemplateLoader struct {
	template   *template.Template
	muTemplate sync.Mutex
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
	loader.muTemplate.Lock()
	defer loader.muTemplate.Unlock()
	if len(bytes) == 0 {
		return fmt.Errorf("empty file")
	}
	if _, err := loader.template.Parse(string(bytes)); err != nil {
		return err
	}
	return nil
}

func (loader *TemplateLoader) Template() *template.Template {
	return loader.template
}
