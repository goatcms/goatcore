package encryptfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestRemove(t *testing.T) {
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
	path := "/mydir1/mydir2/mydir3/mydir4"
	if err := fs.MkdirAll(path, 0777); err != nil {
		t.Error("Fail when create directories", err)
		return
	}
	// test node type
	if fs.Remove("/mydir1/mydir2") == nil {
		t.Error("Remove remove no empty directory is not allowed")
	}
	if err := fs.Remove("/mydir1/mydir2/mydir3/mydir4"); err != nil {
		t.Errorf("Remove should remove empty directory (Error: %v)", err)
	}
	if err := fs.RemoveAll("/mydir1"); err != nil {
		t.Errorf("RemoveAll should remove empty directory (Error: %v )", err)
	}
}
