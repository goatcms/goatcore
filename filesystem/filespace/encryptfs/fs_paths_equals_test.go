package encryptfs

import (
	"io"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestEqualsPath(t *testing.T) {
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
	//Create data
	testData := []byte("There is test data")
	// create directories
	err = fs.WriteFile("mydir1/mydir2/mydir3/myfile.ex", testData, 0777)
	if err != nil {
		t.Error(err)
		return
	}
	reader, err := fs.Reader("/mydir1/mydir2/mydir3/myfile.ex")
	if err != nil {
		t.Error(err)
		return
	}
	buf := make([]byte, 222)
	n, err := reader.Read(buf)
	if err != io.EOF {
		t.Error(err)
		return
	}
	err = reader.Close()
	if err != nil {
		t.Error(err)
		return
	}
	if n != len(testData) {
		t.Errorf("return length should be equal to data size %v %v", n, len(testData))
		return
	}
	if !testbase.ByteArrayEq(buf[:n], testData) {
		t.Error("read data are different ", buf, testData)
	}
}
