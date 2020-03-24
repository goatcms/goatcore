package extcfs

import (
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs/aesgcm256cfs"
)

const (
	// AESGCM256CFS represent AES (Advanced Encryption Standard) with GCM (Galois/Counter Mode)
	AESGCM256CFS = iota
)

// CipherMap contains map to ciper
type CipherMap map[CipherKey]cipherfs.Cipher

// AllCiphers contains map to all default ciphers
var AllCiphers = CipherMap{
	AESGCM256CFS: aesgcm256cfs.NewCipher(),
}
