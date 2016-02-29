package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type DirLoop struct {
	OnFile func(os.FileInfo, string) error
	OnDir  func(os.FileInfo, string) error
	Filter func(os.FileInfo, string) bool
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
		fmt.Println("LoopDir e: %s", dir.Name())

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
			return dirLoopRun(basePath, newSubPath+"/", ctx)
		} else {
			err = ctx.OnFile(dir, newSubPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
