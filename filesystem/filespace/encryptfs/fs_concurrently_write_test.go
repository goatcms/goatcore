package encryptfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/workers"
)

func TestConcurrentlyWrite(t *testing.T) {
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
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		path := randomPath(5)
		for i := workers.MaxJob; i > 0; i-- {
			go writeFileTestHelper(fs, path, t)
		}
	}
}
