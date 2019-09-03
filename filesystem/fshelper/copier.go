package fshelper

import (
	"io"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Copier is helper to copy data bettwent two filesystem
type Copier struct {
	SrcFS    filesystem.Filespace
	SrcPath  string
	DestFS   filesystem.Filespace
	DestPath string
}

// Do clone files from sourcefs to destfs
func (c Copier) Do() (err error) {
	if c.SrcFS.IsFile(c.SrcPath) {
		return c.copyFile()
	}
	return c.copyDirectory()
}

func (c Copier) copyFile() (err error) {
	var (
		reader filesystem.Reader
		writer filesystem.Writer
	)
	if reader, err = c.SrcFS.Reader(c.SrcPath); err != nil {
		return err
	}
	if writer, err = c.DestFS.Writer(c.DestPath); err != nil {
		reader.Close()
		return err
	}
	if _, err = io.Copy(writer, reader); err != nil {
		writer.Close()
		reader.Close()
		return err
	}
	if err = writer.Close(); err != nil {
		reader.Close()
		return err
	}
	if err = reader.Close(); err != nil {
		return err
	}
	return nil
}

func (c Copier) copyDirectory() (err error) {
	var (
		srcFS  filesystem.Filespace
		destFS filesystem.Filespace
	)
	if !c.SrcFS.IsDir(c.SrcPath) {
		return goaterr.Errorf("%s is not a directory", c.SrcPath)
	}
	if srcFS, err = c.SrcFS.Filespace(c.SrcPath); err != nil {
		return err
	}
	if err = c.DestFS.MkdirAll(c.DestPath, filesystem.DefaultUnixDirMode); err != nil {
		return err
	}
	if destFS, err = c.DestFS.Filespace(c.DestPath); err != nil {
		return err
	}
	return Copy(srcFS, destFS, nil)
}
