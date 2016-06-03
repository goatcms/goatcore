package generator

import (
	"fmt"
	"github.com/goatcms/goat-core/varutil"
	"strings"
)

type Generator interface {
	Definitions
}

func (g *Generator) Generate(v Values) (bool, error) {
	modyfied := false
	for index, def := range g.Definitions {
		if _, isset := v[index]; isset {
			continue //omit defined values
		}
		modyfied = true
		if def.Length == 0 {
			def.Length = 12
		}
		switch strings.ToLower(def.Type) {
		case "alpha":
			v[index] = varutil.RandString(def.Length, varutil.AlphaBytes)
		case "numeric":
			v[index] = varutil.RandString(def.Length, varutil.NumericBytes)
		case "alphanumeric":
			v[index] = varutil.RandString(def.Length, varutil.AlphaNumericBytes)
		case "strong":
			v[index] = varutil.RandString(def.Length, varutil.StrongBytes)
		default:
			return true, fmt.Errorf("Unknow type for value for generator ", def)
		}
	}
	return modyfied, nil
}
