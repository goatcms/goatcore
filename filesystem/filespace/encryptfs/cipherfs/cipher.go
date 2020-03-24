package cipherfs

import "github.com/goatcms/goatcore/filesystem"

// Cipher provide encrypt/decrypt functions
type Cipher interface {
	DecryptReader(key []byte, reader filesystem.Reader) (filesystem.Reader, error)
	EncryptWriter(key []byte, writer filesystem.Writer) (filesystem.Writer, error)
	Encrypt(key []byte, data []byte) (encrypted []byte, err error)
	Decrypt(key []byte, data []byte) (decrypted []byte, err error)
}
