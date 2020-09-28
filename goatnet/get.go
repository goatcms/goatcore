package goatnet

import (
	"io/ioutil"
	"net/http"
)

// ReadURL return url content
func ReadURL(url string) (content []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
