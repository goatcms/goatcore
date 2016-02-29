package scaffolding

import (
	"encoding/json"
	"os"
)

type Delimiters struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Config struct {
	Ignore     []string          `json:"ignore"`
	Generate   map[string]string `json:"generate"`
	Delimiters Delimiters        `json:"delimiters"`
}

func readConfig(src string) (*Config, error) {
	c := &Config{}
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&c); err != nil {
		return nil, err
	}
	return c, nil
}
