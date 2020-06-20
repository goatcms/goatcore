package tasks

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
)

func testTaskImplementTaskWriter(ctx app.IOContext, pip pipservices.Pip, broadcast app.Broadcast, closeCB func()) pipservices.TaskWriter {
	return NewTask(ctx, pip, broadcast, closeCB)
}
