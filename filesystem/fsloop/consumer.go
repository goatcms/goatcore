package fsloop

import "github.com/goatcms/goat-core/workers/jobsync"

type Consumer struct {
	lifecycle *jobsync.Lifecycle
	pool      *jobsync.Pool
	loopData  *LoopData
}

func (consumer *Consumer) Loop() {
	defer consumer.pool.Done()
	for {
		// check if anyother job killed job group
		if consumer.lifecycle.IsKilled() {
			return
		}
		// process
		select {
		case row, more := <-consumer.loopData.chans.dirChan:
			if !more {
				return
			}
			if err := consumer.loopData.OnDir(consumer.loopData.Filespace, row); err != nil {
				consumer.lifecycle.Error(err)
				continue
			}
		case row, more := <-consumer.loopData.chans.fileChan:
			if !more {
				return
			}
			if err := consumer.loopData.OnFile(consumer.loopData.Filespace, row); err != nil {
				consumer.lifecycle.Error(err)
				continue
			}
		}
	}
}
