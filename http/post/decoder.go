package post

import (
	"fmt"
	"net/http"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/varutil"
)

// Decoder is global form decoder
type Decoder struct {
	structType types.CustomType
}

// NewDecoder create a form decder instance
func NewDecoder(structType types.CustomType) *Decoder {
	return &Decoder{structType}
}

// Decode tranform request data to structure
func (rd *Decoder) Decode(obj interface{}, req *http.Request) error {
	m := make(map[string]interface{})
	//form
	if err := req.ParseForm(); err != nil {
		return err
	}
	for key, customType := range rd.structType.GetSubTypes() {
		val, ok := req.PostForm[key]
		if !ok {
			return fmt.Errorf("can not find %v key of request form", key)
		}
		data, err := customType.FromString(val[0])
		if err != nil {
			return err
		}
		m[key] = data
		continue
	}
	return varutil.LoadStruct(obj, m, "schema", true)
}
