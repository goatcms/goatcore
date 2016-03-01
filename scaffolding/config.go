package scaffolding

import (
	"encoding/json"
	"os"
)

type Delimiters struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Module struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

type Config struct {
	//Ignore     []string          `json:"ignore"`
	//Generate   map[string]string `json:"generate"`
	Delimiters Delimiters `json:"delimiters"`
	Modules    []Module   `json:"modules"`
}

func (c Config) Read(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&c); err != nil {
		return err
	}
	return nil
}
