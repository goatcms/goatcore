package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
	"os"
	"path/filepath"
)

type Builder struct {
	Src string
	Suffixes     []string
	FileRenderer FileRenderer
	//
	//ResourceName string
	//
	//Templates []string
	//
	//Delimiters Delimiters
	//Data       interface{}
}

func (b *Builder) Init() error {
	if filesystem.IsDir(b.Src) {
		varutil.FixDirPath(&b.Src)
	}
	return nil
}

func (b *Builder) Build(dest string) error {
	if filesystem.IsDir(b.Src) {
		filesystem.MkdirAll(dest)
		varutil.FixDirPath(&dest)
	} else {
		filesystem.MkdirAll(filepath.Dir(dest))
		if err := b.renderFile(b.Src, dest); err != nil {
			return err
		}
		return nil
	}
	loop := filesystem.DirLoop{
		OnFile: b.buildFileFactory(dest),
		OnDir:  b.buildDirFactory(dest),
		Filter: b.filterFactory([]string{
			".git",
		}),
	}
	if err := loop.Run(b.Src); err != nil {
		return err
	}
	return nil
}

func (b *Builder) renderFile(src, dest string) error {
	if !varutil.HasOneSuffix(src, b.Suffixes) {
		return filesystem.CopyFile(src, dest)
	}
	if err := b.FileRenderer.Render(src, dest); err != nil {
		return err
	}
	return nil
}

func (b *Builder) buildFileFactory(dest string) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		return b.renderFile(b.Src+path, dest+path)
	}
}

func (b *Builder) buildDirFactory(dest string) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		return filesystem.MkdirAll(dest + path)
	}
}

func (b *Builder) filterFactory(excludes []string) func(os.FileInfo, string) bool {
	return func(file os.FileInfo, path string) bool {
		return !varutil.IsArrContainStr(excludes, path)
	}
}
