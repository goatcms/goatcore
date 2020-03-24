package extcfs

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/goatcms/goatcore/filesystem"
)

type Buffer struct {
	bytes.Buffer
}

func (Buffer) Close() (err error) {
	return nil
}

func TestEncryptDecryptStory(t *testing.T) {
	var (
		expected  = []byte("Some text")
		key       = []byte("Some Key")
		encrypted []byte
		decrypted []byte
		err       error
	)
	if encrypted, err = NewDefaultCipher().Encrypt(key, expected); err != nil {
		t.Error(err)
		return
	}
	if decrypted, err = NewDefaultCipher().Decrypt(key, encrypted); err != nil {
		t.Error(err)
		return
	}
	if bytes.Compare(expected, decrypted) != 0 {
		t.Errorf("Expected %s and take %v", expected, decrypted)
	}
}

func TestEncryptDecryptStreamStory(t *testing.T) {
	var (
		expected = []byte("Some text")
		key      = []byte("Some Key")
		buff     = &Buffer{}
		writer   filesystem.Writer
		reader   filesystem.Reader
		result   []byte
		err      error
	)
	if writer, err = NewDefaultCipher().EncryptWriter(key, buff); err != nil {
		t.Error(err)
		return
	}
	if _, err = writer.Write(expected); err != nil {
		t.Error(err)
		return
	}
	if writer.Close(); err != nil {
		t.Error(err)
		return
	}
	if reader, err = NewDefaultCipher().DecryptReader(key, buff); err != nil {
		t.Error(err)
		return
	}
	if result, err = ioutil.ReadAll(reader); err != nil {
		t.Error(err)
		return
	}
	if writer.Close(); err != nil {
		t.Error(err)
		return
	}
	if bytes.Compare(expected, result) != 0 {
		t.Errorf("Expected %s and take %v", expected, result)
	}
}
