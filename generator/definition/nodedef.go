package definition

import (
	"fmt"
	"github.com/goatcms/goat-core/generator"
	"reflect"
	"strings"
)

const (
	Separator = ";"
)

type NodeDef struct {
	name    string
	typeDef generator.TypeDef
	nodes   reflect.Value
}

func EncodeNewNodeDef(str string, nodes reflect.Value) (generator.NodeDef, error) {
	params := strings.Split(str, Separator)
	if params[0] == "" {
		return nil, fmt.Errorf("Node must be defined")
	}
	td, err := NewTypeDef(params[1:])
	if err != nil {
		return nil, err
	}
	return NewNodeDef(params[0], td, nodes)
}

func NewNodeDef(name string, typeDef generator.TypeDef, nodes reflect.Value) (generator.NodeDef, error) {
	uValue, err := unpackValue(nodes)
	if err != nil {
		return nil, err
	}
	if nodes.Type().Kind() != reflect.Map {
		return nil, fmt.Errorf("Nodes value must be a map")
	}
	return &NodeDef{
		name:    name,
		typeDef: typeDef,
		nodes:   uValue,
	}, nil
}

func (nd *NodeDef) Name() string {
	return nd.name
}

func (nd *NodeDef) Type() generator.TypeDef {
	return nd.typeDef
}

func (nd *NodeDef) NodesIterator() (generator.NodeDefIterator, error) {
	return NewNodesIterator(nd.nodes)
}
