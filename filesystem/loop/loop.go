package loop

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/workers"
	"github.com/goatcms/goat-core/workers/jobsync"
)

// Loop is a loop on a filespace
type Loop struct {
	scope        app.EventScope
	loopData     *LoopData
	lifecycle    *jobsync.Lifecycle
	consumerPool *jobsync.Pool
}

func NewLoop(loopData *LoopData, scope app.EventScope) *Loop {
	loop := &Loop{
		scope:    scope,
		loopData: loopData,
	}
	loop.loopData.chans.dirChan = make(chan string, ChanSize)
	loop.loopData.chans.fileChan = make(chan string, ChanSize)
	return loop
}

func (loop *Loop) Run(path string) {
	loop.run(path)
}

// Run start process filesystem
func (loop *Loop) run(path string) {
	// lifecycle
	loop.lifecycle = jobsync.NewLifecycle(workers.DefaultTimeout, true)
	if loop.scope != nil {
		loop.scope.On(app.ErrorEvent, loop.KillSlot)
		loop.scope.On(app.KillEvent, loop.KillSlot)
	}
	// producer
	producerPool := jobsync.NewPool(workers.MaxJob)
	producer := &Producer{
		lifecycle: loop.lifecycle,
		pool:      producerPool,
		loopData:  loop.loopData,
		path:      path,
	}
	producerCounter := producerPool.Add(1)
	if producerCounter != 1 {
		panic("filesystem.Loop it is not possible to start producer job")
	}
	go producer.Loop()
	// consumer
	loop.consumerPool = jobsync.NewPool(workers.MaxJob)
	consumer := &Consumer{
		lifecycle: loop.lifecycle,
		pool:      loop.consumerPool,
		loopData:  loop.loopData,
	}
	consumerCounter := loop.consumerPool.Add(workers.MaxJob)
	for i := 0; i < consumerCounter; i++ {
		go consumer.Loop()
	}
	// lifecycle
	go func() {
		producerPool.Wait()
		close(loop.loopData.chans.dirChan)
		close(loop.loopData.chans.fileChan)
	}()
}

// Wait wait for job finish
func (loop *Loop) Wait() {
	loop.consumerPool.Wait()
}

func (loop *Loop) KillSlot(interface{}) error {
	loop.lifecycle.Kill()
	return nil
}

func (loop *Loop) Errors() []error {
	return loop.lifecycle.Errors()
}
