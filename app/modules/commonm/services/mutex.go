package services

const (
	//LockRW lock for read and write
	LockRW = true
	//LockR lock for read only
	LockR = false
)

// LockMap describe lock resources
// It is important. The group can be difined only once.
// It must be unlock after finish.
type LockMap map[string]bool

// SharedMutex lock global resources by names
type SharedMutex interface {
	Lock(resources LockMap) (handler UnlockHandler)
}

// UnlockHandler lock app resources
type UnlockHandler interface {
	Unlock()
}
