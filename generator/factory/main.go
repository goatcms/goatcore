package generator

import (
	"fmt"
	"github.com/goatcms/goat-core/generator"
	"github.com/goatcms/goat-core/varutil"
	"math/rand"
	"strconv"
	"strings"
)

const (
	String = "string"
)

type Factory struct {
}

func NewFactory() (generator.DataFactory, error) {
	return generator.DataFactory(&Factory{}), nil
}

func (f *Factory) BuildInt(def TypeDef) (string, error) {
	switch strings.ToLower(def.GeneratorName()) {
	case "rand":
		return rand.New(time.Now().UnixNano()).Int()
	case "const":
		args := def.Args()
		if len(args) < 3 {
			return "", fmt.Errorf("Value (third param) for const int generator are required ", def)
		}
		i, err := strconv.Atoi(args[2])
		if err != nil {
			return "", fmt.Errorf("can not convert value to string", args[2], def)
		}
		return i
	default:
		return "", fmt.Errorf("Unknow type for value for generator ", def)
	}
}

func (f *Factory) BuildString(def TypeDef) (string, error) {
	lengthAttr := def.Params().Key("length")
	length, err := strconv.Atoi(lengthAttr)
	if err != nil {
		length = 12
	}
	switch strings.ToLower(def.GeneratorName()) {
	case "alpha":
		return varutil.RandString(length, varutil.AlphaBytes)
	case "numeric":
		return varutil.RandString(length, varutil.NumericBytes)
	case "alphanumeric":
		return varutil.RandString(length, varutil.AlphaNumericBytes)
	case "strong":
		return varutil.RandString(length, varutil.StrongBytes)
	case "const":
		args := def.Args()
		if len(args) < 3 {
			return "", fmt.Errorf("Value (third param) for const string generator are required ", def)
		}
		return args[2]
	default:
		return "", fmt.Errorf("Unknow type for value for generator ", def)
	}
}
