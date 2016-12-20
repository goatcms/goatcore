package loop

import (
	"github.com/goatcms/goat-core/filesystem"
)

type Loop struct {
	onFile filesystem.LoopOn
	onDir  filesystem.LoopOn
	filter filesystem.LoopFilter
}

func NewLoop() filesystem.Loop {
	return filesystem.Loop(&Loop{
		onFile: nil,
		onDir:  nil,
		filter: nil,
	})
}

func (l *Loop) OnFile(cb filesystem.LoopOn) {
	l.onFile = cb
}

func (l *Loop) OnDir(cb filesystem.LoopOn) {
	l.onDir = cb
}

func (l *Loop) Filter(cb filesystem.LoopFilter) {
	l.filter = cb
}

func (l *Loop) Run(fs filesystem.Filespace) error {
	return l.runLoop(fs, "")
}

func (l *Loop) runLoop(fs filesystem.Filespace, subPath string) error {
	list, err := fs.ReadDir(subPath)
	if err != nil {
		return err
	}
	for _, dir := range list {
		if dir.Name() == "." || dir.Name() == ".." {
			continue
		}
		newSubPath := subPath + dir.Name()
		if l.Filter != nil && !l.filter(fs, newSubPath, dir) {
			continue
		}
		if dir.IsDir() {
			if l.onDir != nil {
				if err = l.onDir(fs, newSubPath, dir); err != nil {
					return err
				}
			}
			if err = l.runLoop(fs, newSubPath+"/"); err != nil {
				return err
			}
		} else {
			if err = l.onFile(fs, newSubPath, dir); err != nil {
				return err
			}
		}
	}
	return nil
}
