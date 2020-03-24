package encryptfs

import (
	"os"
	"testing"
	"time"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestModTime(t *testing.T) {
	var (
		basefs, fs filesystem.Filespace
		err        error
		fileInfo   os.FileInfo
		modTime    time.Time
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
	// Prepare test data
	testData := []byte("There is test data")
	path := "/mydir1/mydir2/mydir3/myfile.ex"
	fs.WriteFile(path, testData, 0777)
	// test data
	if fileInfo, err = fs.Lstat(path); err != nil {
		t.Error(err)
	}
	if modTime = fileInfo.ModTime(); modTime.IsZero() {
		t.Errorf("Time %v can not be unix zero date", modTime)
	}
	if modTime != fileInfo.ModTime() {
		t.Errorf("Modification time should be modified only on change")
	}
	fs.WriteFile(path, testData, 0777)
	if modTime == fileInfo.ModTime() {
		t.Errorf("Modification time should be updated on change")
	}
}
