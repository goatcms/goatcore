package abstracttype

import (
	"testing"

	"github.com/goatcms/goat-core/filesystem/filespace/memfs"
	"github.com/goatcms/goat-core/testbase"
	"github.com/goatcms/goat-core/testbase/mocks/mockmultipart"
	"github.com/goatcms/goat-core/types"
)

const (
	// TestFilePath is file path for filesystem test
	TestFilePath = "path/to/file.ex"
	// TestFileContent is a test file content
	TestFileContent = "1234567890qwertyuiopasdfghjkl;'zxcvbnm,./!@#$%^&*()"
)

func TestFilespaceConverter_FromString(t *testing.T) {
	memFilespace, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	converter := NewFilespaceConverter(memFilespace)
	iInterface, err := converter.FromString(TestFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	file, ok := iInterface.(types.File)
	if !ok {
		t.Errorf("converter.FromString result is not types.File instance %v", file)
		return
	}
	if file.Path() != TestFilePath {
		t.Errorf("Path of file must be equels %v != %v", TestFilePath, file.Path())
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
	fh := mockmultipart.NewFileHeader([]byte(TestFileContent))
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
	if !testbase.ByteArrayEq(fileContent, []byte(TestFileContent)) {
		t.Errorf("converter.FromString result is not types.File instance \n%v \n !=\n %v", string(fileContent), TestFileContent)
		return
	}
}

func TestFilespaceConverter_ToString(t *testing.T) {
	memFilespace, err := memfs.NewFilespace()
	if err != nil {
		t.Error(err)
		return
	}
	file := NewFile(memFilespace, TestFilePath)
	converter := NewFilespaceConverter(memFilespace)
	str, err := converter.ToString(file)
	if err != nil {
		t.Error(err)
		return
	}
	if str != TestFilePath {
		t.Errorf("converter.ToString result is incorrect: %v", str)
		return
	}
}
