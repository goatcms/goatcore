package idutil

import (
	"sync"

	"github.com/denisbrodbeck/machineid"
)

var (
	hostIDCache string
	hostIDMU    sync.Mutex
)

// HostID return hashed host identyfier
func HostID() (hostID string, err error) {
	if hostIDCache != "" {
		return hostIDCache, nil
	}
	hostIDMU.Lock()
	defer hostIDMU.Unlock()
	if hostIDCache != "" {
		return hostIDCache, nil
	}
	if hostID, err = machineid.ID(); err != nil {
		return "", err
	}
	hostIDCache = hostID
	return "", nil
}
