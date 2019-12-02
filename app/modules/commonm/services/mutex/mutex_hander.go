package mutex

import (
	"github.com/goatcms/goatcore/app/modules/commonm/services"
)

type mutexRow struct {
	Name  string
	Value bool
}

type unlockHandler struct {
	list        []mutexRow
	sharedMutex *SharedMutex
}

// Unlock locked resources
func (hander *unlockHandler) Unlock() {
	for _, row := range hander.list {
		mu := hander.sharedMutex.get(row.Name)
		if row.Value == services.LockR {
			mu.RUnlock()
		} else {
			mu.Unlock()
		}
	}
}
