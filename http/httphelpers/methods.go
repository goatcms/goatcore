package httphelpers

import (
	"mime"
	"net/http"
)

// ParseForm parse http form
func ParseForm(req *http.Request) (err error) {
	if IsMultipart(req) {
		if err = req.ParseMultipartForm(111); err != nil {
			return err
		}
	} else {
		if err = req.ParseForm(); err != nil {
			return err
		}
	}
	return nil
}

// IsMultipart check if request is a multipart media type
func IsMultipart(req *http.Request) bool {
	v := req.Header.Get("Content-Type")
	if v == "" {
		return false
	}
	d, _, err := mime.ParseMediaType(v)
	if err != nil || d != "multipart/form-data" {
		return false
	}
	return true
}
