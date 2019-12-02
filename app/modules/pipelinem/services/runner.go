package services

import (
	"io"

	"github.com/goatcms/goatcore/app"
)

// Pip describe single commands pipeline
type Pip struct {
	Name string
	Body io.Reader
}

// PipRunner run command pipeline
type PipRunner interface {
	Run(ctx app.IOContext) (err error)
}
