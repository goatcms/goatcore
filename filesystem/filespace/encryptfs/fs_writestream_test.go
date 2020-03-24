package encryptfs

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/extcfs"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
	"github.com/goatcms/goatcore/testbase"
)

func TestWriteStreamAndRead(t *testing.T) {
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
	path := "/mydir1/mydir2/mydir3/myfile.ex"
	writer, err := fs.Writer(path)
	if err != nil {
		t.Error(err)
		return
	}
	n, err := writer.Write(testData)
	if err != nil {
		t.Error(err)
		return
	}
	if n != len(testData) {
		t.Errorf("return length should be equal to data size %v %v", n, len(testData))
		return
	}
	err = writer.Close()
	if err != nil {
		t.Error(err)
		return
	}
	readData, err := fs.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	if !testbase.ByteArrayEq(readData, testData) {
		t.Error("read data are different ", readData, testData)
	}
}
