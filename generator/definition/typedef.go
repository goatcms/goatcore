package definition

import (
	"fmt"
	"github.com/goatcms/goat-core/generator"
)

type TypeDef struct {
	typeName      string
	generatorName string
	params        Params
}

func NewTypeDef(a []string) (*TypeDef, error) {
	td := TypeDef{}
	if a[0] == "" {
		return nil, fmt.Errorf("Type name can not be null")
	}
	if a[1] == "" {
		return nil, fmt.Errorf("Generator name can not be null")
	}
	td.typeName = a[0]
	td.generatorName = a[1]
	td.params = Params(a[2:])
	return &td, nil
}

func (t *TypeDef) TypeName() string {
	return t.typeName
}

func (t *TypeDef) GeneratorName() string {
	return t.generatorName
}

func (t *TypeDef) Params() generator.TypeDefParams {
	return t.params
}
