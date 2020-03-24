package aesgcm256cfs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs"
)

// Cipher provide encrypt/decrypt functions
type Cipher struct{}

// NewCipher Create new NewCipher Instance
func NewCipher() cipherfs.Cipher {
	return instance
}

// DecryptReader create decrypt stream for AES GCM
func (Cipher) DecryptReader(key []byte, stream filesystem.Reader) (filesystem.Reader, error) {
	return newReader(key, stream)
}

// EncryptWriter create encrypt stream for AES GCM
func (Cipher) EncryptWriter(key []byte, stream filesystem.Writer) (filesystem.Writer, error) {
	return newWriter(key, stream), nil
}

// Encrypt AES GCM data with key
func (Cipher) Encrypt(key []byte, data []byte) (encrypted []byte, err error) {
	var (
		block cipher.Block
		gcm   cipher.AEAD
		nonce []byte
	)
	if key, err = hash256(key); err != nil {
		return nil, err
	}
	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	if gcm, err = cipher.NewGCM(block); err != nil {
		return nil, err
	}
	nonce = make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Decrypt AES GCM data with key
func (Cipher) Decrypt(key []byte, data []byte) (decrypted []byte, err error) {
	var (
		block cipher.Block
		gcm   cipher.AEAD
		nonce []byte
	)
	if key, err = hash256(key); err != nil {
		return nil, err
	}
	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	if gcm, err = cipher.NewGCM(block); err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, data = data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, data, nil)
}
