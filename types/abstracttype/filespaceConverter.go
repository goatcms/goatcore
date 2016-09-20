package abstracttype

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/varutil"
)

// FilespaceConverter is converter for strings
type FilespaceConverter struct {
	fs filesystem.Filespace
}

// NewFilespaceConverter create new file converter for filespace
func NewFilespaceConverter(fs filesystem.Filespace) *FilespaceConverter {
	return &FilespaceConverter{fs}
}

func (fc *FilespaceConverter) persistFile(reader io.Reader) (string, error) {
	n1 := strconv.FormatInt(time.Now().Unix(), 16)
	n2 := strconv.FormatInt(time.Now().UnixNano(), 16)
	n3 := varutil.RandString(6, varutil.AlphaNumericBytes)
	name := n1 + "." + n2 + "." + n3
	wr, err := fc.fs.Writer(name)
	if err != nil {
		return name, err
	}
	defer wr.Close()
	_, err = io.Copy(wr, reader)
	return name, err
}

// FromString decode string value
func (fc *FilespaceConverter) FromString(path string) (interface{}, error) {
	return NewFile(fc.fs, path), nil
}

// FromMultipart convert multipartdata to string
func (fc *FilespaceConverter) FromMultipart(fh types.FileHeader) (interface{}, error) {
	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	path, err := fc.persistFile(file)
	if err != nil {
		return nil, err
	}
	return NewFile(fc.fs, path), nil
}

// ToString change object to string
func (fc *FilespaceConverter) ToString(ival interface{}) (string, error) {
	typesFile, ok := ival.(types.File)
	if !ok {
		return "", fmt.Errorf("FilespaceConverter.ToString aceppt only types.File as input")
	}
	return typesFile.Path(), nil
}
