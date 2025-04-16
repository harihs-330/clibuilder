package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Help     HelpInfo          `yaml:"help"`
	Commands map[string]string `yaml:"commands"`
}

type HelpInfo struct {
	Description       string            `yaml:"description"`
	Examples          []string          `yaml:"examples"`
	Usage             []string          `yaml:"usage"`
	AvailableCommands map[string]string `yaml:"available_commands"`
	Flags             []string          `yaml:"flags"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
