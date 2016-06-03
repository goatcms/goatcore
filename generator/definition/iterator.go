package definition

import (
	"fmt"
	"github.com/goatcms/goat-core/generator"
	"reflect"
)

type Iterator struct {
	keys []reflect.Value
	i    int
	val  reflect.Value
}

func NewNodesIterator(vmap reflect.Value) (generator.NodeDefIterator, error) {
	vmap, err := unpackValue(vmap)
	if err != nil {
		return nil, err
	}
	return generator.NodeDefIterator(&Iterator{
		keys: vmap.MapKeys(),
		i:    0,
		val:  vmap,
	}), nil
}

func (l *Iterator) HasNext() bool {
	return l.i < len(l.keys)
}

func (l *Iterator) Next() (generator.NodeDef, error) {
	subNodeKey := l.keys[l.i]
	subNodeValue, err := unpackValue(l.val.MapIndex(subNodeKey))
	if err != nil {
		return nil, err
	}
	l.i++
	if subNodeValue.IsNil() {
		return EncodeNewNodeDef(subNodeKey.String(), reflect.ValueOf(map[string]string{}))
	}
	switch subNodeValue.Type().Kind() {
	case reflect.Map:
		return EncodeNewNodeDef(subNodeKey.String(), subNodeValue)
	default:
		return nil, fmt.Errorf("Unsupported type ", subNodeValue.Type().Kind())
	}
}
