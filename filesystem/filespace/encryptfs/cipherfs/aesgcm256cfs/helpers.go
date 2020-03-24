package aesgcm256cfs

import "golang.org/x/crypto/sha3"

// hash256 create 256bit hash
func hash256(in []byte) (out []byte, err error) {
	var hash = sha3.New256()
	if _, err = hash.Write(in); err != nil {
		return out, err
	}
	return hash.Sum(nil), nil
}
