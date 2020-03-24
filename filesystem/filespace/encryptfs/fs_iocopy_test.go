package encryptfs

import (
	"io"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestIOCopy(t *testing.T) {
	var (
		err        error
		basefs, fs filesystem.Filespace
		writer     filesystem.Writer
		reader     filesystem.Reader
		content    []byte
	)
	t.Parallel()
	if basefs, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
	}
	if fs, err = NewEncryptFS(basefs, Settings{
		Salt:     []byte("salt"),
		Secret:   []byte("secret"),
		HostOnly: false,
		Cipher:   extcfs.NewDefaultCipher(),
	}); err != nil {
		t.Error(err)
	}
	if err = fs.WriteFile("file.txt", []byte("file text"), 0766); err != nil {
		t.Error(err)
		return
	}
	if reader, err = fs.Reader("file.txt"); err != nil {
		t.Error(err)
		return
	}
	if writer, err = fs.Writer("file2.txt"); err != nil {
		t.Error(err)
		return
	}
	if _, err = io.Copy(writer, reader); err != nil {
		t.Error(err)
		return
	}
	reader.Close()
	writer.Close()
	if content, err = fs.ReadFile("file2.txt"); err != nil {
		t.Error(err)
		return
	}
	if string(content) != "file text" {
		t.Error(err)
		return
	}
}
