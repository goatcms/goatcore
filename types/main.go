package types

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
)

// ValidErrors represent a list of validation messages
type ValidErrors []string

// CustomType represent a list of validation messages
type CustomType interface {
	SQLType() string
	HTMLType() string
	HasAttribute(name string) bool

	IsValid(string) (ValidErrors, error)
}

// Validator represent a single validator
type Validator func(string) (string, error)
