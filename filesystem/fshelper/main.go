package fshelper

import (
	"io"

	"github.com/goatcms/goatcore/filesystem"
)

// StreamCopy copy file from source filesystem to destination filesystem with a stream loop
func StreamCopy(sourcefs, destfs filesystem.Filespace, subPath string) (err error) {
	var (
		reader filesystem.Reader
		writer filesystem.Writer
	)
	if reader, err = sourcefs.Reader(subPath); err != nil {
		return err
	}
	if writer, err = destfs.Writer(subPath); err != nil {
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
