package fsloop

import "github.com/goatcms/goat-core/filesystem"

const (
	ChanSize        = 500
	MinExtraJobData = 500
)

type Chans struct {
	dirChan  chan string
	fileChan chan string
}

type LoopData struct {
	chans      Chans
	Filespace  filesystem.Filespace
	FileFilter filesystem.LoopFilter
	DirFilter  filesystem.LoopFilter
	OnFile     filesystem.LoopOn
	OnDir      filesystem.LoopOn
}
