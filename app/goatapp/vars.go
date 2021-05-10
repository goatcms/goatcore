package goatapp

import "fmt"

const (
	// ConfigFilePath is path to main config file
	ConfigFilePath = "/config/config_{{env}}.json"
	// CWDPath is default Current Working Directory (CWD) path
	CWDPath = "./"
)

var (
	ErrAppNameIsRequired = fmt.Errorf("Application name is required")
)
