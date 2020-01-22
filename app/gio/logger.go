package gio

import (
	"fmt"
	"time"

	"github.com/goatcms/goatcore/app"
)

// Logger write outupt with timestamp
type Logger struct {
	out app.Output
	cid string
}

// NewLogger return new Logger instance
func NewLogger(out app.Output, cid string) app.Output {
	return &Logger{
		out: out,
		cid: cid,
	}
}

// Writer is the interface that wraps the basic Write method.
func (logger *Logger) Write(p []byte) (n int, err error) {
	formatted := logger.format(string(p))
	if n, err = logger.out.Write([]byte(formatted)); err != nil {
		return n, err
	}
	return len(p), err
}

// Printf print to multiple outputs.
func (logger *Logger) Printf(format string, a ...interface{}) (err error) {
	formatted := logger.format(format)
	return logger.out.Printf(formatted)
}

func (logger *Logger) format(s string) (result string) {
	timeStr := time.Now().Format("2006-01-02 15:04:05.000000000")
	return fmt.Sprintf(" [%s] %s : %s\n", timeStr, logger.cid, s)
}
