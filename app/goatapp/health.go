package goatapp

import (
	"fmt"
	"sync"

	"github.com/goatcms/goatcore/app"
)

// HealthCheckers is implements app.HealthCheckers
type HealthCheckers struct {
	mu   sync.RWMutex
	data map[string]app.HealthCheckerCallback
}

// NewHealthChecker create new instance of app.HealthChecker
func NewHealthCheckers() (healths app.AppHealthCheckers) {
	return newHealthCheckers()
}

func newHealthCheckers() (healths *HealthCheckers) {
	return &HealthCheckers{
		data: make(map[string]app.HealthCheckerCallback),
	}
}

// HealthChecker return health's names
func (healths *HealthCheckers) HealthCheckerNames() (keys []string) {
	healths.mu.RLock()
	keys = make([]string, len(healths.data))
	i := 0
	for key := range healths.data {
		keys[i] = key
		i++
	}
	healths.mu.RUnlock()
	return
}

// HealthChecker return a health by name
func (healths *HealthCheckers) HealthChecker(name string) (health app.HealthCheckerCallback) {
	healths.mu.RLock()
	health = healths.data[name]
	healths.mu.RUnlock()
	return
}

// SetHealthChecker set new comamnd
func (healths *HealthCheckers) SetHealthChecker(name string, health app.HealthCheckerCallback) (err error) {
	healths.mu.Lock()
	defer healths.mu.Unlock()
	if _, ok := healths.data[name]; ok {
		return fmt.Errorf("HealthChecker %s is defined twice", name)
	}
	healths.data[name] = health
	return
}
