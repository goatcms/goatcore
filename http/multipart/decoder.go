package multipart

import (
	"fmt"
	"net/http"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/varutil"
)

const _24M = (1 << 20) * 24

// Decoder is global form decoder
type Decoder struct {
	structType types.StructType
}

// NewDecoder create a form decder instance
func NewDecoder(structType types.StructType) *Decoder {
	return &Decoder{structType}
}

// Decode tranform request data to structure
func (rd *Decoder) Decode(obj interface{}, req *http.Request) error {
	m := make(map[string]interface{})
	//form
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := req.ParseMultipartForm(0); err != nil {
		return err
	}
	for key, customType := range rd.structType {
		val, ok := req.PostForm[key]
		if ok {
			data, err := customType.FromString(val[0])
			if err != nil {
				return err
			}
			m[key] = data
			continue
		}
		fha, ok := req.MultipartForm.File[key]
		if !ok || len(fha) < 1 {
			return fmt.Errorf("can not find %v key of request form", key)
		}
		fh := fha[0]
		data, err := customType.FromMultipart(fh)
		if err != nil {
			return err
		}
		m[key] = data
		continue
	}
	return varutil.LoadStruct(obj, m, "schema", true)
}
