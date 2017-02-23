package smtpmail

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type TestConfig struct {
	SenderConfig Config `json:"sender"`
	FromAddress  string `json:"fromAddress"`
	ToAddress    string `json:"toAddress"`
}

func LoadTestConfig() (*TestConfig, error) {
	config := &TestConfig{}

	path, err := filepath.Abs("../../tests/smtp.json")
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
