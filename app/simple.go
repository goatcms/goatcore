package app

// TypeConverter convert from type to other (string->int etc)
type TypeConverter func(interface{}) (interface{}, error)

// EventCallback is a callback function with data
type EventCallback func(interface{}) error

// Callback is a callback function
type Callback func() error
