package loop

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/workers"
	"github.com/goatcms/goat-core/workers/paraller"
)

// Loop is a loop on a filespace
type Loop struct {
	FS     filesystem.Filespace
	Scope  app.EventScope
	OnFile filesystem.LoopOn
	OnDir  filesystem.LoopOn
	Filter filesystem.LoopFilter

	chans       *chans
	producerJob workers.Job
	consumerJob workers.Job
}

// KillSlot is a slot to kill the loop
func (l *Loop) KillSlot(interface{}) error {
	l.producerJob.Kill()
	l.consumerJob.Kill()
	return nil
}

// Run start process filesystem
func (l *Loop) Run(path string) error {
	l.chans = &chans{
		dirChan:  make(chan string, 500),
		fileChan: make(chan string, 500),
		baseChan: make(chan string, 500),
	}
	l.chans.baseChan <- path
	l.producerJob = paraller.NewParaller(producerBody{
		fs:     l.FS,
		filter: l.Filter,
		chans:  l.chans,
	})
	l.producerJob.Defer(l.closeChans)
	l.consumerJob = paraller.NewParaller(consumerBody{
		fs:     l.FS,
		onFile: l.OnFile,
		onDir:  l.OnDir,
		chans:  l.chans,
	})
	if l.Scope != nil {
		l.Scope.On(app.ErrorEvent, l.producerJob.KillSlot)
		l.Scope.On(app.KillEvent, l.consumerJob.KillSlot)
	}
	if err := l.producerJob.Run(); err != nil {
		return err
	}
	if err := l.consumerJob.Run(); err != nil {
		return err
	}
	return nil
}

// Wait wait for job finish
func (l *Loop) closeChans() error {
	close(l.chans.baseChan)
	close(l.chans.dirChan)
	close(l.chans.fileChan)
	return nil
}

// Wait wait for job finish
func (l *Loop) Wait() error {
	if err := l.producerJob.Wait(); err != nil {
		return err
	}
	if err := l.consumerJob.Wait(); err != nil {
		return err
	}
	return nil
}
