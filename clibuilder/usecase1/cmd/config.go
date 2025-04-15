package cmd

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Commands []string `json:"commands"`
	Paths    []string `json:"paths"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
