package scaffolding

import (
	"encoding/json"
	"fmt"
	"github.com/goatcms/goat-core/varutil"
	"io/ioutil"
	"os"
	"strings"
)

type Generator struct {
	Definitions Definitions
}

func NewGenerator() *Generator {
	g := &Generator{
		Definitions: Definitions{},
	}
	return g
}

func (g *Generator) LoadDefinitions(path string) error {
	file, err := os.Open(path + generatorDefSuffix)
	if err != nil {
		g.Definitions = map[string]GeneratorDefinition{}
		return nil
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&g.Definitions); err != nil {
		return err
	}
	return nil
}

func (g *Generator) LoadValues(path string) error {
	file, err := os.Open(path + generatorGenSuffix)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&g.Values); err != nil {
		return err
	}
	return nil
}

func (g *Generator) GenerateValues() error {
	for index, def := range g.Definitions {
		if _, isset := g.Values[index]; isset {
			continue //omit defined values
		}
		if def.Length == 0 {
			def.Length = 12
		}
		switch strings.ToLower(def.Type) {
		case "alpha":
			g.Values[index] = varutil.RandString(def.Length, varutil.AlphaBytes)
		case "numeric":
			g.Values[index] = varutil.RandString(def.Length, varutil.NumericBytes)
		case "alphanumeric":
			g.Values[index] = varutil.RandString(def.Length, varutil.AlphaNumericBytes)
		case "strong":
			g.Values[index] = varutil.RandString(def.Length, varutil.StrongBytes)
		default:
			return fmt.Errorf("Unknow type for value for generator ", def)
		}
	}
	return nil
}

func (g *Generator) SaveValues(path string) error {
	b, err := json.Marshal(g.Values)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path+generatorGenSuffix, b, 0777)
	if err != nil {
		return err
	}
	return nil
}
