package i18mem

import (
	"fmt"
	"sync"

	"github.com/goatcms/goatcore/i18n"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// I18Mem contains translation storaged in memory
type I18Mem struct {
	muTranlsates sync.RWMutex
	translates   map[string]string
}

// NewI18N return new I18Mem instance
func NewI18N() i18n.I18N {
	return NewI18Mem()
}

// NewI18Mem return new I18Mem instance
func NewI18Mem() *I18Mem {
	return &I18Mem{
		translates: map[string]string{},
	}
}

// Set add new translations
func (i18 *I18Mem) Set(values map[string]string) {
	i18.muTranlsates.Lock()
	defer i18.muTranlsates.Unlock()
	for key, value := range values {
		i18.translates[key] = value
	}
}

// SetDefault add new translations
func (i18 *I18Mem) SetDefault(values map[string]string) {
	i18.muTranlsates.Lock()
	defer i18.muTranlsates.Unlock()
	for key, value := range values {
		if _, ok := i18.translates[key]; !ok {
			i18.translates[key] = value
		}
	}
}

// Translate return translate for key
func (i18 *I18Mem) Translate(key string, values ...interface{}) (string, error) {
	i18.muTranlsates.RLock()
	defer i18.muTranlsates.RUnlock()
	format, ok := i18.translates[key]
	if !ok {
		return "", goaterr.Errorf("Unknown translate for %s key", key)
	}
	return fmt.Sprintf(format, values...), nil
}
