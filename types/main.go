package types

import "mime/multipart"

const (
	// NoErr is default string for no error response
	NoErr = ""
	// Unique is unique attribute
	Unique = "unique"
	// Primary is primary key attribute for sql
	Primary = "primary"
	// NotNull is not null attribute for sql
	NotNull = "notnull"
	// Required is not null and check length>0 attribute for sql
	Required = "required"
	// SQLSeparator represent separator for SQL
	SQLSeparator = "_"
)

// FileHeader is multipart interface
type FileHeader interface {
	Open() (multipart.File, error)
}

// MessageMap contains object errors
type MessageMap interface {
	Get(key string) MessageList
	GetAll() map[string]MessageList
	Add(key, msg string)
}

// MessageList represent list of field errors
type MessageList interface {
	GetAll() []string
	Add(msgkey string)
}

// Validator represent a single validator
type Validator func(interface{}, string, MessageMap) error

// CustomType represent type interface
type CustomType interface {
	SingleCustomType
	GetSubTypes() map[string]CustomType
	AddSubTypes(string, map[string]CustomType)
}

// SingleCustomType represent one type interface
type SingleCustomType interface {
	MetaType
	TypeConverter
	TypeValidator
}

// MetaType represent type data
type MetaType interface {
	SQLType() string
	HTMLType() string
	HasAttribute(name string) bool
	GetAttribute(name string) string
}

// TypeConverter convert type
type TypeConverter interface {
	FromString(string) (interface{}, error)
	FromMultipart(fh FileHeader) (interface{}, error)
	ToString(interface{}) (string, error)
}

// TypeValidator valid type data
type TypeValidator interface {
	Valid(interface{}, string, MessageMap) error
}
