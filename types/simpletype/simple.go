package simpletype

import "github.com/goatcms/goat-core/types"

// SimpleCustomType is abstract for simple types
type SimpleCustomType struct {
	SQLTypeName  string
	HTMLTypeName string
	Validators   []types.Validator
	Attributes   []string
}

// SQLType return a sql type name
func (s *SimpleCustomType) SQLType() string {
	return s.SQLTypeName
}

// HTMLType return a HTML type name
func (s *SimpleCustomType) HTMLType() string {
	return s.HTMLTypeName
}

// HasAttribute return true if type has attribute
func (s *SimpleCustomType) HasAttribute(name string) bool {
	for _, v := range s.Attributes {
		if name == v {
			return true
		}
	}
	return false
}

// IsValid chuck a value is valid
func (s *SimpleCustomType) IsValid(value string) (types.ValidErrors, error) {
	var errs []string
	for _, validator := range s.Validators {
		errStr, err := validator(value)
		if err != nil {
			return nil, err
		}
		if errStr != types.NoErr {
			errs = append(errs, errStr)
		}
	}
	return errs, nil
}
