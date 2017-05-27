package dbdriver

import (
	"fmt"
	"sync"

	"github.com/goatcms/goatcore/db"
)

var (
	driversMu sync.RWMutex
	drivers   map[string]db.Driver
)

func init() {
	drivers = make(map[string]db.Driver)
}

func Register(name string, engine db.Driver) error {
	driversMu.Lock()
	defer driversMu.Unlock()
	if _, ok := drivers[name]; ok {
		return fmt.Errorf("dbdrive: Register called twice for driver %s", name)
	}
	drivers[name] = engine
	return nil
}

func Driver(name string) (db.Driver, error) {
	driversMu.RLock()
	defer driversMu.RUnlock()
	engine, ok := drivers[name]
	if !ok {
		return nil, fmt.Errorf("dbdriver: %s driver undefined", name)
	}
	return engine, nil
}

// Drivers returns a (unsorted) list of the names of the registered drivers.
func Drivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()
	list := make([]string, len(drivers))
	i := 0
	for name := range drivers {
		list = append(list, name)
		i++
	}
	return list
}
