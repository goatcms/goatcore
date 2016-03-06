package generator

import (
	"github.com/goatcms/goat-core/varutil"
)

type GeneratorDefinition struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
}

type Definitions map[string]GeneratorDefinition

func (d *Definitions) Read(path string) error {
	return varutil.ReadJson(path, d)
}
