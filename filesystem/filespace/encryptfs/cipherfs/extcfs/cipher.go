package extcfs

import (
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Cipher provide encrypt/decrypt functions
type Cipher struct {
	mapping         CipherMap
	defaultCiper    cipherfs.Cipher
	defaultCiperKey CipherKey
}

// NewDefaultCipher Create new Cipher with default data
func NewDefaultCipher() (result cipherfs.Cipher) {
	var err error
	if result, err = NewCipher(AESGCM256CFS, AllCiphers); err != nil {
		panic(err)
	}
	return result
}

// NewCipher Create new Cipher Instance
func NewCipher(defaultCiperKey CipherKey, mapping CipherMap) (result cipherfs.Cipher, err error) {
	c := Cipher{
		mapping:         mapping,
		defaultCiperKey: defaultCiperKey,
	}
	if c.defaultCiper = mapping[defaultCiperKey]; c.defaultCiper == nil {
		return nil, goaterr.Errorf("map %v must contains default cipher %v", mapping, defaultCiperKey)
	}
	return c, nil
}

// DecryptReader create decrypt stream for AES GCM
func (c Cipher) DecryptReader(key []byte, stream filesystem.Reader) (reader filesystem.Reader, err error) {
	var (
		ckey       CipherKey
		p          = make([]byte, 4)
		fileCipher cipherfs.Cipher
	)
	if _, err = stream.Read(p); err != nil {
		return nil, err
	}
	ckey = NewCipherKey(p)
	if fileCipher = c.mapping[ckey]; fileCipher == nil {
		return nil, goaterr.Errorf("Unknow cipher for %v key", ckey)
	}
	return fileCipher.DecryptReader(key, stream)
}

// EncryptWriter create encrypt stream for AES GCM
func (c Cipher) EncryptWriter(key []byte, stream filesystem.Writer) (writer filesystem.Writer, err error) {
	if _, err = stream.Write(c.defaultCiperKey.ToBinary()); err != nil {
		return nil, err
	}
	return c.defaultCiper.EncryptWriter(key, stream)
}

// Encrypt AES GCM data with key
func (c Cipher) Encrypt(key []byte, data []byte) (encrypted []byte, err error) {
	if data, err = c.defaultCiper.Encrypt(key, data); err != nil {
		return nil, err
	}
	return append(c.defaultCiperKey.ToBinary(), data...), nil
}

// Decrypt AES GCM data with key
func (c Cipher) Decrypt(key []byte, data []byte) (decrypted []byte, err error) {
	var (
		ckey       = NewCipherKey(data[:4])
		fileCipher cipherfs.Cipher
	)
	if fileCipher = c.mapping[ckey]; fileCipher == nil {
		return nil, goaterr.Errorf("Unknow cipher for %v key", ckey)
	}
	return fileCipher.Decrypt(key, data[4:])
}
