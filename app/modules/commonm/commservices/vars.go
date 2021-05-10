package commservices

import "time"

// Info provide application details
type Info interface {
	// Name is provided by app instance
	// Version is provided by app instance
	Authors() []InfoPerson
	Description() string
	License() InfoLicense
	PoweredBy() string
}

// InfoPerson represent person
type InfoPerson interface {
	Name() string
	Email() string
}

// InfoLicense provide license info
type InfoLicense interface {
	Company() string
	End() time.Time
	Start() time.Time
	Type() string
	URL() string
}
