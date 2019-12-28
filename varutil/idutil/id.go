package idutil

import "strconv"

// ID return unique id
func ID() int64 {
	idMU.Lock()
	defer idMU.Unlock()
	id++
	return id
}

// StringID return unique string id
func StringID() string {
	return strconv.FormatInt(ID(), 36)
}
