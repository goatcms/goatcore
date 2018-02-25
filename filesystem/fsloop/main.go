package fsloop

import (
	"sync"

	"github.com/goatcms/goatcore/filesystem"
)

const (
	// ChanSize is default channel size
	ChanSize = 1000
	// StepClose is close step id
	StepClose = 999
)

// Chans contains channels and mutex for loop
type Chans struct {
	muDirChan  sync.Mutex
	dirChan    chan string
	muFileChan sync.Mutex
	fileChan   chan string
}

// LoopData is loop data container
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
