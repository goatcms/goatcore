package workers

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/goatcms/goatcore/filesystem"
)

// ResponseContext is a error container for current context
type ResponseContext struct {
	mu          sync.Mutex
	out         string
	errors      []error
	wg          sync.WaitGroup
	closed      bool
	stream      filesystem.Writer
	id          string
	taskName    string
	triggeredBy string

	fs   filesystem.Filespace
	path string

	tasksLogs map[string]string
}

// NewResponseContext create a new error context
func NewResponseContext(fs filesystem.Filespace, taskName, triggeredBy string) (responseContext *ResponseContext, err error) {
	var (
		writer filesystem.Writer
		now    = time.Now()
		nowStr = DateTimeStr(now)
	)
	id := NewResponseID(taskName)
	path := taskName + "/" + id
	if err = fs.MkdirAll(filepath.Dir(path), filesystem.DefaultUnixDirMode); err != nil {
		return nil, err
	}
	if writer, err = fs.Writer(path + ".stream.log"); err != nil {
		return nil, err
	}
	if err = fs.MkdirAll(taskName, filesystem.DefaultUnixDirMode); err != nil {
		return nil, err
	}
	responseContext = &ResponseContext{
		id:          id,
		errors:      []error{},
		fs:          fs,
		path:        path,
		tasksLogs:   map[string]string{},
		stream:      writer,
		closed:      false,
		taskName:    taskName,
		triggeredBy: triggeredBy,
	}
	responseContext.Print("MAIN", fmt.Sprintf("Triggered by %v", triggeredBy))
	responseContext.Print("MAIN", fmt.Sprintf("START TIME %v", nowStr))
	return responseContext, nil
}

// Print add task output to a output stream
func (ctx *ResponseContext) Print(taskName, out string) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.preventClose()
	ctx.print(taskName, out)
}

func (ctx *ResponseContext) print(taskName, out string) {
	var (
		now    = time.Now()
		nowStr = TimeStr(now)
	)
	if out == "" {
		return
	}
	wdata := fmt.Sprintf("*%s %s: %s\n", nowStr, taskName, out)
	ctx.tasksLogs[taskName] += out + "\n"
	ctx.stream.Write([]byte(wdata))
	ctx.out += wdata
}

// ID return response ID
func (ctx *ResponseContext) ID() string {
	return ctx.id
}

// Output return context output
func (ctx *ResponseContext) Output() string {
	return ctx.out
}

// AddError add new error to error list
func (ctx *ResponseContext) AddError(taskName string, err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.preventClose()
	ctx.errors = append(ctx.errors, err)
	ctx.print(taskName, fmt.Sprintf("ERROR %s", err.Error()))
}

// Errors return errors
func (ctx *ResponseContext) Errors() []error {
	return ctx.errors
}

// Close output and persist logs
func (ctx *ResponseContext) Close() (err error) {
	var (
		now    = time.Now()
		nowStr = DateTimeStr(now)
	)
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.print("MAIN", fmt.Sprintf("FINISH TIME %s", nowStr))
	ctx.closed = true
	if err = ctx.stream.Close(); err != nil {
		return err
	}
	raport := ctx.prepareRaport()
	if len(ctx.errors) == 0 {
		ctx.fs.WriteFile(ctx.path+".success.log", []byte(raport), filesystem.DefaultUnixFileMode)
	} else {
		ctx.fs.WriteFile(ctx.path+".fail.log", []byte(raport), filesystem.DefaultUnixFileMode)
	}
	return nil
}

// IsClosed check id response was closed
func (ctx *ResponseContext) IsClosed() bool {
	return ctx.closed
}

// WaitGroup return response *sync.WaitGroup
func (ctx *ResponseContext) WaitGroup() *sync.WaitGroup {
	return &ctx.wg
}

func (ctx *ResponseContext) preventClose() {
	if ctx.closed {
		panic("can not access to response after close")
	}
}

func (ctx *ResponseContext) prepareRaport() string {
	var (
		out string
	)
	if len(ctx.errors) == 0 {
		out += fmt.Sprintf("Process success\nTriggered by %v", ctx.triggeredBy)
	} else {
		out += fmt.Sprintf("Process failed\nTriggered by %v", ctx.triggeredBy)
	}
	for taskName, content := range ctx.tasksLogs {
		out += fmt.Sprintf("\n\n##### %v:\n%v", taskName, content)
	}
	return out
}
