package fsloop

import (
	"os"

	"github.com/goatcms/goatcore/workers/jobsync"
)

type Producer struct {
	lifecycle *jobsync.Lifecycle
	pool      *jobsync.Pool
	loopData  *LoopData
	path      string
}

func (producer *Producer) Loop() {
	defer producer.pool.Done()
	readDir, err := producer.loopData.Filespace.ReadDir(producer.path)
	if err != nil {
		producer.lifecycle.Error(err)
		return
	}
	producer.processList(producer.path, readDir)
}

func (producer *Producer) processList(basePath string, readDir []os.FileInfo) bool {
	for _, node := range readDir {
		if node.Name() == "." || node.Name() == ".." {
			continue
		}
		nodePath := basePath + node.Name()
		if node.IsDir() {
			if producer.loopData.DirFilter != nil {
				if producer.loopData.DirFilter(producer.loopData.Filespace, nodePath) {
					if producer.loopData.OnDir != nil {
						producer.loopData.chans.dirChan <- nodePath
					}
					if isKilled := producer.processDir(nodePath); isKilled {
						return true
					}
				}
				continue
			}
			if producer.loopData.OnDir != nil {
				producer.loopData.chans.dirChan <- nodePath
			}
			if isKilled := producer.processDir(nodePath); isKilled {
				return true
			}
		} else {
			if producer.loopData.OnFile == nil {
				continue
			}
			if producer.loopData.FileFilter != nil && !producer.loopData.FileFilter(producer.loopData.Filespace, nodePath) {
				continue
			}
			producer.processFile(nodePath)
		}
		if producer.lifecycle.IsKilled() {
			return true
		}
	}
	return false
}

func (producer *Producer) processFile(nodePath string) {
	producer.loopData.chans.fileChan <- nodePath
}

func (producer *Producer) processDir(nodePath string) bool {
	jobCounter := producer.pool.Add(1)
	if jobCounter == 0 {
		// recursive process if no free thread
		readDir, err := producer.loopData.Filespace.ReadDir(nodePath)
		if err != nil {
			producer.lifecycle.Error(err)
			return true
		}
		basePath := nodePath + "/"
		producer.processList(basePath, readDir)
		return false
	}
	// create new job
	newProducer := &Producer{
		lifecycle: producer.lifecycle,
		pool:      producer.pool,
		loopData:  producer.loopData,
		path:      nodePath + "/",
	}
	go newProducer.Loop()
	return false
}
