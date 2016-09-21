package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/testbase"
	"github.com/goatcms/goat-core/testbase/mocks/mockmultipart"
	"github.com/goatcms/goat-core/types"
)

const (
	TEST_FILE_PATH    = "path/to/file.ex"
	TEST_FILE_CONTENT = "1234567890qwertyuiopasdfghjkl;'zxcvbnm,./!@#$%^&*()"
)

func TestFilespaceConverter_FromString(t *testing.T) {
	memFilespace, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	converter := NewFilespaceConverter(memFilespace)
	iInterface, err := converter.FromString(TEST_FILE_PATH)
	if err != nil {
		t.Error(err)
		return
	}
	file, ok := iInterface.(types.File)
	if !ok {
		t.Errorf("converter.FromString result is not types.File instance %v", file)
		return
	}
	if file.Path() != TEST_FILE_PATH {
		t.Errorf("Path of file must be equels %v != %v", TEST_FILE_PATH, file.Path())
		return
	}
}

func TestFilespaceConverter_FromMultipart(t *testing.T) {
	memFilespace, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	converter := NewFilespaceConverter(memFilespace)
	fh := mockmultipart.NewFileHeader([]byte(TEST_FILE_CONTENT))
	iInterface, err := converter.FromMultipart(fh)
	if err != nil {
		t.Error(err)
		return
	}
	file, ok := iInterface.(types.File)
	if !ok {
		t.Errorf("converter.FromString result is not types.File instance %v", file)
		return
	}
	fileContent, err := file.Filespace().ReadFile(file.Path())
	if err != nil {
		t.Error(err)
		return
	}
	if !testbase.ByteArrayEq(fileContent, []byte(TEST_FILE_CONTENT)) {
		t.Errorf("converter.FromString result is not types.File instance \n%v \n !=\n %v", string(fileContent), TEST_FILE_CONTENT)
		return
	}
}

func TestFilespaceConverter_ToString(t *testing.T) {
	memFilespace, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	file := NewFile(memFilespace, TEST_FILE_PATH)
	converter := NewFilespaceConverter(memFilespace)
	str, err := converter.ToString(file)
	if err != nil {
		t.Error(err)
		return
	}
	if str != TEST_FILE_PATH {
		t.Errorf("converter.ToString result is incorrect: %v", str)
		return
	}
}
