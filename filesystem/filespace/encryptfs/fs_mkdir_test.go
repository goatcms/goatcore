package encryptfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestMkdir(t *testing.T) {
	var (
		basefs, fs filesystem.Filespace
		err        error
	)
	t.Parallel()
	// init
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
	// create directories
	path := "/mydir1/mydir2/mydir3"
	if err := fs.MkdirAll(path, 0777); err != nil {
		t.Error("Fail when create directories", err)
	}
	// test node type
	if !fs.IsDir("/mydir1/mydir2") {
		t.Error("node is not a directory or not exists")
	}
	if !fs.IsDir(path) {
		t.Error("node is not a directory or not exists")
	}
	if fs.IsDir("/noExistPath") {
		t.Error("node is not a directory or not exists")
	}
}
