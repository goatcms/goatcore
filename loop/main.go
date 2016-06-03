package loop

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
}
