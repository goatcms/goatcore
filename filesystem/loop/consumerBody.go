package loop

import (
	"runtime"

	"github.com/goatcms/goat-core/filesystem"
)

type consumerBody struct {
	fs     filesystem.Filespace
	onFile filesystem.LoopOn
	onDir  filesystem.LoopOn
	chans  *chans
}

func (j consumerBody) Step() (bool, error) {
	select {
	case row, more := <-j.chans.dirChan:
		if !more {
			return false, nil
		}
		if err := j.onDir(j.fs, row); err != nil {
			return false, err
		}
	case row, more := <-j.chans.fileChan:
		if !more {
			return false, nil
		}
		if err := j.onFile(j.fs, row); err != nil {
			return false, err
		}
	default:
		runtime.Gosched()
	}
	return true, nil
}
