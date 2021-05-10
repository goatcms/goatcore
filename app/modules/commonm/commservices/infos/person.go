package infos

import (
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
)

type PersonParams struct {
	Name  string
	Email string
}

// Person SandboxesPerson is a tool to menage sandboxes.
type Person struct {
	params PersonParams
}

func NewPerson(params PersonParams) commservices.InfoPerson {
	return &Person{
		params: params,
	}
}

// Name return firsname and lastname or/and nickname
func (person *Person) Name() string {
	return person.params.Name
}

// Email return person email
func (person *Person) Email() string {
	return person.params.Email
}
