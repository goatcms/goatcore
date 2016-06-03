package definition

import (
	//"fmt"
	"github.com/goatcms/goat-core/generator"
	"github.com/goatcms/goat-core/testbase"
	"reflect"
	"testing"
)

func TestIterator(t *testing.T) {
	defObject := map[string]interface{}{
		"key1;number;const;11":       map[string]string{},
		"key2;string;generatorname1": interface{}(nil),
		"key3;string;generatorname2": nil,
		"key4;map;mapgeneratorname": map[string]interface{}{
			"subkey1;string;const;value": nil,
		},
	}

	outData := map[string]NodeDef{
		"key1": NodeDef{
			name: "key1",
			typeDef: generator.TypeDef(&TypeDef{
				typeName:      "number",
				generatorName: "const",
				params:        Params{"11"},
			}),
			nodes: reflect.ValueOf(map[string]interface{}{}),
		},
		"key2": NodeDef{
			name: "key2",
			typeDef: generator.TypeDef(&TypeDef{
				typeName:      "string",
				generatorName: "generatorname1",
				params:        Params{},
			}),
			nodes: reflect.ValueOf(map[string]interface{}{}),
		},
		"key3": NodeDef{
			name: "key3",
			typeDef: generator.TypeDef(&TypeDef{
				typeName:      "string",
				generatorName: "generatorname2",
				params:        Params{},
			}),
			nodes: reflect.ValueOf(map[string]interface{}{}),
		},
		"key4": NodeDef{
			name: "key4",
			typeDef: generator.TypeDef(&TypeDef{
				typeName:      "map",
				generatorName: "mapgeneratorname",
				params:        Params{},
			}),
			nodes: reflect.ValueOf(map[string]interface{}{
				"subkey1": NodeDef{
					name: "subkey1",
					typeDef: generator.TypeDef(&TypeDef{
						typeName:      "string",
						generatorName: "const",
						params:        Params{"value"},
					}),
					nodes: reflect.ValueOf(map[string]interface{}{}),
				},
			}),
		},
	}

	iterator, err := NewNodesIterator(reflect.ValueOf(defObject))
	if err != nil {
		t.Error(err)
	}

	counter := 0
	for iterator.HasNext() {
		row, err := iterator.Next()
		if err != nil {
			t.Error(err)
			return
		}
		expectOut := outData[row.Name()]
		isEquals, err := compareNodeDef(row, generator.NodeDef(&expectOut))
		if !isEquals {
			t.Error("records are incorrect ", row, expectOut)
			return
		}
		counter++
	}

	if counter != len(outData) {
		t.Error("iterator run ", counter, " times, expected ", len(outData), " times")
	}
}

func compareIterator(iterator1 generator.NodeDefIterator, iterator2 generator.NodeDefIterator) (bool, error) {
	for iterator1.HasNext() {
		if !iterator2.HasNext() {
			return false, nil
		}
		e1, err := iterator1.Next()
		if err != nil {
			return false, err
		}
		e2, err := iterator2.Next()
		if err != nil {
			return false, err
		}
		isEquals, err := compareNodeDef(e1, e2)
		if err != nil {
			return false, err
		}
		if !isEquals {
			return false, nil
		}
	}
	if iterator2.HasNext() {
		return false, nil
	}
	return true, nil
}

func compareNodeDef(o1 generator.NodeDef, o2 generator.NodeDef) (bool, error) {
	if o1.Name() != o2.Name() {
		return false, nil
	}
	//iterate
	t1 := o1.Type()
	t2 := o2.Type()
	if t1.TypeName() != t1.TypeName() || t1.GeneratorName() != t1.GeneratorName() {
		return false, nil
	}
	isEquals, err := testbase.ComparePublicEquals(t1.Params(), t2.Params())
	if err != nil {
		return false, err
	}
	if !isEquals {
		return false, nil
	}
	//iterate
	iterator1, err := o1.NodesIterator()
	if err != nil {
		return false, err
	}
	iterator2, err := o1.NodesIterator()
	if err != nil {
		return false, err
	}
	return compareIterator(iterator1, iterator2)
}
