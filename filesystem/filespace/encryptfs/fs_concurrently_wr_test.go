package encryptfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/workers"
)

func TestConcurrentlyRead(t *testing.T) {
	var (
		basefs, fs  filesystem.Filespace
		err         error
		expetedText = "some data"
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
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		rPath := randomPath(5)
		if err = fs.WriteFile(rPath, []byte(expetedText), filesystem.DefaultUnixDirMode); err != nil {
			t.Error(err)
			return
		}
		for i := workers.MaxJob; i > 0; i-- {
			go (func() {
				var data []byte
				if data, err = fs.ReadFile(rPath); err != nil {
					t.Error(err)
					return
				}
				if string(data) != expetedText {
					t.Errorf("Expected '%s' text and take %s", expetedText, data)
					return
				}
			})()
		}
	}
}
