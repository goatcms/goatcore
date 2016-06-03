package generator

const (
	MapNode    = "map"
	ArrayNode  = "array"
	ObjectNode = "object"
	StringNode = "string"
	IntNode    = "int"
)

type TypeDef interface {
	TypeName() string
	GeneratorName() string
	Params() TypeDefParams
}

type TypeDefParams interface {
	Args() []string
	Key(name string) string
}

type NodeDef interface {
	Name() string
	Type() TypeDef
	NodesIterator() (NodeDefIterator, error)
}

type NodeDefIterator interface {
	HasNext() bool
	Next() (NodeDef, error)
}

type DataFactory interface {
	/*BuildArray(def TypeDef) ([]interface{}, error)
	BuildObject(def TypeDef) (interface{}, error)
	BuildMap(def TypeDef)  (interface{}, error)*/
	BuildInt(def TypeDef)  (int, error)
	BuildString(def TypeDef) (string, error)
}

type Builder interface {
	Def(def NodeDef)
	//Generator(gen DataGenerator)
	//Data(data Data)
	Run()
}
