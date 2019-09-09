package memfs

import (
	"io"
	"testing"

	"github.com/goatcms/goatcore/testbase"
)

func TestEqualsPath(t *testing.T) {
	t.Parallel()
	// init
	fs, err := NewFilespace()
	if err != nil {
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
