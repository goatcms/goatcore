package readbuilder

import (
	generator "github.com/goatcms/goat-core/generator"
	r "github.com/goatcms/goat-core/varutil/r"
	"reflect"
	"strings"
)

type Builder struct {
	nodeDef generator.NodeDef
	data    interface{}
}

func NewBuilder(nodeDef generator.NodeDef, facotry generator.DataFactory, data interface{}) (*Builder, error) {
	return &Builder{
		nodeDef: nodeDef,
		data:    data,
		facotry: factory,
	}, nil
}

func (b *Builder) Run() (interface{}, error) {
	tn := b.nodeDef.Type().TypeName()
	value := reflect.ValueOf(b.data)
	uValue := r.UnpackValue(value)
	dtn := uValue.Type().String()
	if b.data != nil && tn != dtn {
		return nil, fmt.Errorf("Incompatible data " + uValue + " and definition type " + tn)
	}
	switch dtn {
	case "string":
		return runString(value)
	case "int":
		return runInt(value)
	case "[]interface {}":
		return runArray(value)
	case "map[string]interface {}":
		return runMap(value)
	default:
		return nil, fmt.Errorf("Unknow type " + uValue)
	}
}

func (b *Builder) runString(rvalue reflect.Value, nodeDef generator.NodeDef) (interface{}, error) {
	var (
		value interface{}
		scan  string
		err   error
	)
	if !data.IsNil() {
		value = r.UnpackValue(rvalue).Interface()
	} else {
		value, err = b.facotry.BuildString(nodeDef.Type())
		if err != nil {
			return nil, err
		}
	}
	fmt.Print("[", value, "]:")
	fmt.Scanf(scan)
	if scan != "" {
		value = interface{}(scan)
	}
	return value, nil
}

func (b *Builder) runInt(rvalue reflect.Value, nodeDef generator.NodeDef) (interface{}, error) {
	var (
		value interface{}
		ival  int
		scan  string
		err   error
	)
	if !data.IsNil() {
		value = r.UnpackValue(rvalue).Interface()
	} else {
		value, err = b.facotry.BuildInt(nodeDef.Type())
		if err != nil {
			return nil, err
		}
	}
	for {
		fmt.Print("[", value, "]:")
		fmt.Scanf(scan)
		if scan == "" {
			break
		}
		if ival, err = strconv.Atoi(scan); err != nil {
			fmt.Printf("Incorrect number.\n")
			continue
		}
		value = interface{}(ival)
	}
	return value, nil
}

func (b *Builder) runArray(rvalue reflect.Value, nodeDef generator.NodeDef) (interface{}, error) {
	var (
		value []interface{}
		scan  string
		err   error
	)
	if !data.IsNil() {
		value = r.UnpackValue(rvalue).Interface().([]interface{})
	} else {
		value = []interface{}{}
	}
	for {
		fmt.Println("your array: ", value)
		fmt.Print("[", value, "]:")
		fmt.Scanf(scan)
		if scan == "" {
			break
		}
		if ival, err = strconv.Atoi(scan); err != nil {
			fmt.Printf("Incorrect number.\n")
			continue
		}
		value = interface{}(ival)
	}
	return value, nil
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
