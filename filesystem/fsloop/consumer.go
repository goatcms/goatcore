package fsloop

import (
	"runtime"

	"github.com/goatcms/goatcore/workers/jobsync"
)

// Consumer is object represent single loop consumer task
type Consumer struct {
	lifecycle *jobsync.Lifecycle
	pool      *jobsync.Pool
	loopData  *LoopData
}

// Loop start data process (use "go Loop()" to run it in new goroutine)
func (consumer *Consumer) Loop() {
	defer consumer.pool.Done()
	for {
		if consumer.lifecycle.IsKilled() {
			return
		}
		if len(consumer.loopData.chans.dirChan) == 0 &&
			len(consumer.loopData.chans.fileChan) == 0 {
			if consumer.lifecycle.Step() == StepClose {
				return
			}
			runtime.Gosched()
		} else {
			if len(consumer.loopData.chans.dirChan) != 0 {
				select {
				case row, more := <-consumer.loopData.chans.dirChan:
					if !more {
						continue
					}
					if err := consumer.loopData.OnDir(consumer.loopData.Filespace, row); err != nil {
						consumer.lifecycle.Error(err)
					}
				default:
					continue
				}
			}
			if len(consumer.loopData.chans.fileChan) != 0 {
				select {
				case row, more := <-consumer.loopData.chans.fileChan:
					if !more {
						continue
					}
					if err := consumer.loopData.OnFile(consumer.loopData.Filespace, row); err != nil {
						consumer.lifecycle.Error(err)
					}
				default:
					continue
				}
			}
		}
	}
}
