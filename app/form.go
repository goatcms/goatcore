package app

import (
	"github.com/goatcms/goatcore/messages"
)

// Form represent a form data
type Form interface {
	Valid() (messages.MessageMap, error)
	Data() interface{}
}
