package gtprovider

import (
	"fmt"
	"text/template"
	"sync"

	"github.com/goatcms/goatcore/filesystem"
)

// TemplateLoader provide method to load templates from filesystem
type TemplateLoader struct {
	template   *template.Template
	muTemplate sync.Mutex
}

// NewTemplateLoader create TemplateLoader instance
func NewTemplateLoader(template *template.Template) *TemplateLoader {
	return &TemplateLoader{
		template: template,
	}
}

// Load get all templates code form files in subPath and add it to template
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
		return fmt.Errorf("%v: %v", subPath, err)
	}
	return nil
}

// Template return loaded template
func (loader *TemplateLoader) Template() *template.Template {
	return loader.template
}
