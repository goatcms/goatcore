package fsloop

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// Loop is a loop on a filespace
type Loop struct {
	scope        app.EventScope
	loopData     *LoopData
	lifecycle    *jobsync.Lifecycle
	consumerPool *jobsync.Pool
}

// NewLoop create new loop instance
func NewLoop(loopData *LoopData, scope app.EventScope) *Loop {
	loop := &Loop{
		scope:    scope,
		loopData: loopData,
	}
	loop.loopData.chans.dirChan = make(chan string, ChanSize)
	loop.loopData.chans.fileChan = make(chan string, ChanSize)
	return loop
}

// Run start process filesystem
func (loop *Loop) Run(path string) {
	if path == "" {
		path = "./"
	}
	// lifecycle
	loop.lifecycle = jobsync.NewLifecycle(workers.DefaultTimeout, true)
	if loop.scope != nil {
		loop.scope.On(app.ErrorEvent, loop.KillSlot)
		loop.scope.On(app.KillEvent, loop.KillSlot)
	}
	// producer
	producentMaxJob := loop.loopData.Producents
	if producentMaxJob == 0 || producentMaxJob > workers.MaxJob {
		producentMaxJob = workers.MaxJob
	}
	producerPool := jobsync.NewPool(producentMaxJob)
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
	consumerMaxJob := loop.loopData.Consumers
	if consumerMaxJob == 0 || consumerMaxJob > workers.MaxJob {
		consumerMaxJob = workers.MaxJob
	}
	loop.consumerPool = jobsync.NewPool(consumerMaxJob)
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
		loop.lifecycle.NextStep(StepClose)
		close(loop.loopData.chans.dirChan)
		close(loop.loopData.chans.fileChan)
	}()
}

// Wait wait for job finish
func (loop *Loop) Wait() {
	loop.consumerPool.Wait()
}

// KillSlot is function implement event listeenr. It is use to kill lifecycle on scope kill and error event.
func (loop *Loop) KillSlot(interface{}) error {
	loop.lifecycle.Kill()
	return nil
}

// Errors return loop errors
func (loop *Loop) Errors() []error {
	return loop.lifecycle.Errors()
}
