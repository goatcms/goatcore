package fsloop

import (
	"sync"

	"github.com/goatcms/goatcore/filesystem"
)

const (
	ChanSize = 1000

	StepClose = 1
)

type Chans struct {
	muDirChan  sync.Mutex
	dirChan    chan string
	muFileChan sync.Mutex
	fileChan   chan string
}

type LoopData struct {
	chans      Chans
	Filespace  filesystem.Filespace
	FileFilter filesystem.LoopFilter
	DirFilter  filesystem.LoopFilter
	OnFile     filesystem.LoopOn
	OnDir      filesystem.LoopOn
	Consumers  int
	Producents int
}
