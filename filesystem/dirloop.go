package filesystem

import (
	"io/ioutil"
	"os"
	"strings"
)

type LoopOn func(os.FileInfo, string) error
type LoopFilter func(os.FileInfo, string) bool

type DirLoop struct {
	OnFile LoopOn
	OnDir  LoopOn
	Filter LoopFilter
}

func (ctx *DirLoop) Run(path string) error {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return dirLoopRun(path, "", ctx)
}

func dirLoopRun(basePath, subPath string, ctx *DirLoop) error {
	list, err := ioutil.ReadDir(basePath + subPath)
	if err != nil {
		return err
	}

	for _, dir := range list {
		if dir.Name() == "." || dir.Name() == ".." {
			continue
		}

		newSubPath := subPath + dir.Name()
		if ctx.Filter != nil && !ctx.Filter(dir, newSubPath) {
			continue
		}

		if dir.IsDir() {
			if ctx.OnDir != nil {
				if err = ctx.OnDir(dir, newSubPath); err != nil {
					return err
				}
			}
			err = dirLoopRun(basePath, newSubPath+"/", ctx)
			if err != nil {
				return err
			}
		} else {
			err = ctx.OnFile(dir, newSubPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
