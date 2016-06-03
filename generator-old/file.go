package generator

import (
	"github.com/goatcms/goat-core/filesystem"
)

const (
	DefSuffix = ".def.json" // use for data definition
	GenSuffix = ".gen.json" // use to save generated data
	DocSuffix = ".doc.json" // use for module documentation
)

type GeneratedResource struct {
	valuePath       string
	definitionsPath string
	generator       Generator
	Values          Values
}

func NewGeneratedResource(p string) (*GeneratedResource, error) {
	r := &GeneratedResource{
		valuePath:       p + GenSuffix,
		definitionsPath: p + DefSuffix,
		generator:       Generator{},
		Values:          Values{},
	}
	if filesystem.IsFile(r.definitionsPath) {
		if err := r.generator.Read(r.definitionsPath); err != nil {
			return nil, err
		}
	}
	if filesystem.IsFile(r.valuePath) {
		if err := r.Values.Read(r.valuePath); err != nil {
			return nil, err
		}
	}
	modyfied, err := r.generator.Generate(r.Values)
	if err != nil {
		return nil, err
	}
	if modyfied {
		if err := r.Values.Write(r.valuePath); err != nil {
			return nil, err
		}
	}
	return r, nil
}
