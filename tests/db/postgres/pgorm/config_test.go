package pgorm

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type TestConfig struct {
	URL string `json:"url"`
}

var (
	config *TestConfig
)

func LoadTestConfig() (*TestConfig, error) {
	if config != nil {
		return config, nil
	}
	config = &TestConfig{}
	path, err := filepath.Abs("../../../db_postgres.json")
	if err != nil {
		return nil, err
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, config); err != nil {
		return nil, err
	}
	return config, nil
}
