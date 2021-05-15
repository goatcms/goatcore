package tasks

import (
	"testing"

	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

func TestTaskImplementTaskWriter(t *testing.T) {
	var _ pipservices.TaskWriter = &Task{}
}
