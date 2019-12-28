package idutil

import "sync"

var (
	idMU sync.Mutex
	id   int64 = 1
)
