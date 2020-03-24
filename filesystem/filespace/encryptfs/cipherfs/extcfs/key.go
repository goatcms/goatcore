package extcfs

import (
	"encoding/binary"
)

// CipherKey is default cipher key type
type CipherKey uint32

// NewCipherKey create new cipher key from bytes
func NewCipherKey(b []byte) CipherKey {
	return CipherKey(binary.LittleEndian.Uint32(b))
}

// ToBinary return binary key representation
func (key CipherKey) ToBinary() (b []byte) {
	b = make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(key))
	return b
}
