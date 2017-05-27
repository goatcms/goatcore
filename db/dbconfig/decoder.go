package dbconfig

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcore/varutil/totype"
)

// Decoder provide decoder to configuration string
type Decoder struct {
	plainmap map[string]string
}

// NewDecoderFromKeyValueString create new decoder with data from db string
func NewDecoderFromKeyValueString(s string) (*Decoder, error) {
	result := make(map[string]string)
	s = strings.Replace(s, "\t", "", -1)
	sections := strings.Split(s, " ")
	for _, v := range sections {
		if v == "" {
			continue
		}
		kdata := strings.Split(v, "=")
		if len(kdata) < 2 {
			return nil, fmt.Errorf("Unknow key for %s", v)
		}
		if len(kdata) > 2 {
			return nil, fmt.Errorf("Too many equal sign %s", v)
		}
		result[kdata[0]] = kdata[1]
	}
	return &Decoder{
		plainmap: result,
	}, nil
}

//Get return string key value
func (d *Decoder) Get(key, alternative string) string {
	value, ok := d.plainmap[key]
	if !ok {
		return alternative
	}
	return value
}

// GetInt64 return int64 key value
func (d *Decoder) GetInt64(key string, alternative int64) int64 {
	svalue, ok := d.plainmap[key]
	if !ok {
		return alternative
	}
	value, err := totype.StringToInt64(svalue)
	if err != nil {
		return alternative
	}
	return value
}

// GetInt return int key value
func (d *Decoder) GetInt(key string, alternative int) int {
	svalue, ok := d.plainmap[key]
	if !ok {
		return alternative
	}
	value, err := totype.StringToInt(svalue)
	if err != nil {
		return alternative
	}
	return value
}
