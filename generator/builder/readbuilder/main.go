package readbuilder

import (
	"github.com/goatcms/goat-core/generator"
	"github.com/goatcms/goat-core/varutil/r"
	"reflect"
	"strings"
)

type Builder struct {
	nodeDef generator.NodeDef
	data 		interfce{}
}

func NewBuilder(nodeDef generator.NodeDef, facotry generator.DataFactory, data interfce{}) (*Builder, error) {
	return &Builder{
		nodeDef: nodeDef,
		data: data,
		facotry: factory,
	}, nil
}

func (b *Builder) Run() (interface{}, error) {
	tn := b.nodeDef.Type().TypeName()
	dv := r.UnpackValue(reflect.ValueOf(b.data))
	dtn := dv.Type().Name()
	if b.data != nil && tn != dtn {
		return nil, fmt.Errorf("Incompatible data "+dt+" and definition type "+tn)
	}
	switch(dtn) {
	case "string":
		return runString()
	default:
		return nil, fmt.Errorf("Unknow type "+dtn)
	}
}

func (b *Builder) runString(data interface{}, nodeDef generator.NodeDef) (interface{}, error) {
	if data != nil &&  {
		// generate
	}
	return data, nil
}

/*
func build(nodeDef generator.NodeDef, val reflect.Value) (interface{}, error) {
	if err := isValidType(nodeDef, val); err != nil {
		return nil, err
	}
	switch nodeDef.Type().TypeName() {
	case generator.MapNode:
		var asd =
	case generator.ArrayNode:
	case generator.StringNode:
	case generator.IntNode:
	}
}

func getOrCreate(nodeDef generator.NodeDef, fromObject reflect.Value) error {
	switch nodeDef.Type().TypeName() {
	case generator.MapNode:
		return fmt.Errorf("NodeDef describe MapNode and get ", nodeDef.Type().TypeName())
	case generator.ArrayNode:
		return fmt.Errorf("NodeDef describe ArrayNode and get ", nodeDef.Type().TypeName())
	/*case generator.ObjectNode: - unsupported
		return fmt.Errorf("NodeDef describe ObjectNode and get ", nodeDef.Type().TypeName())* /
	case generator.StringNode:
		return fmt.Errorf("NodeDef describe StringNode and get ", nodeDef.Type().TypeName())
	case generator.IntNode:
		return fmt.Errorf("NodeDef describe IntNode and get ", nodeDef.Type().TypeName())
	}
	return nil
}

func isValidType(nodeDef generator.NodeDef, val reflect.Value) error {
	switch nodeDef.Type().Kind() {
	case generator.MapNode:
		return fmt.Errorf("NodeDef describe MapNode and get ", nodeDef.Type().TypeName())
	case generator.ArrayNode:
		return fmt.Errorf("NodeDef describe ArrayNode and get ", nodeDef.Type().TypeName())
	case generator.ObjectNode:
		return fmt.Errorf("NodeDef describe ObjectNode and get ", nodeDef.Type().TypeName())
	case generator.StringNode:
		return fmt.Errorf("NodeDef describe StringNode and get ", nodeDef.Type().TypeName())
	case generator.IntNode:
		return fmt.Errorf("NodeDef describe IntNode and get ", nodeDef.Type().TypeName())
	}
	return nil
}
*/
