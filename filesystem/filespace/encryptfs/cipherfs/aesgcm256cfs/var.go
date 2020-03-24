package aesgcm256cfs

import "github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs"

// Instance is Cipher global instance
var instance cipherfs.Cipher = &Cipher{}
