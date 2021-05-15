package idutil

import (
	"encoding/binary"
	"strconv"

	"github.com/denisbrodbeck/machineid"
	"golang.org/x/crypto/sha3"
)

var (
	hostID, obfuscateHostID, correlationHostID = initValues()
)

// HostID return hashed host identyfier
func HostID() (hostID string) {
	return hostID
}

func ObfuscateHostID() []byte {
	return obfuscateHostID
}

func CorrelationHostID() string {
	return correlationHostID
}

func initValues() (hostID string, obfuscate []byte, correlation string) {
	var (
		err error
	)
	// prepare hostID
	if hostID, err = machineid.ID(); err != nil {
		panic(err)
	}
	// obfuscate
	var hash = sha3.New256()
	if _, err = hash.Write([]byte(hostID[:len(hostID)/2-2])); err != nil {
		panic(err)
	}
	if _, err = hash.Write(defaultObfuscateKey); err != nil {
		panic(err)
	}
	if _, err = hash.Write([]byte(hostID[len(hostID)/2+2:])); err != nil {
		panic(err)
	}
	obfuscate = hash.Sum(nil)
	// prepare host id
	unit := binary.BigEndian.Uint64([]byte{obfuscate[3], obfuscate[5], obfuscate[7], obfuscate[10], obfuscate[12], obfuscate[14], obfuscate[16], obfuscate[18]})
	correlation = strconv.FormatUint(unit, 36)
	return
}

// HostHash256 create 256bit hash
func HostHash256(key []byte) (out []byte, err error) {
	var hash = sha3.New256()
	if _, err = hash.Write([]byte(hostID)); err != nil {
		return
	}
	if _, err = hash.Write(key); err != nil {
		return
	}
	return hash.Sum(nil), nil
}
