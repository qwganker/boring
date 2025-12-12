package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type WebConfig struct {
	BasicAuthUsers map[string]string `yaml:"basic_auth_users"`
}

func LoadWebConfig(filename string) (*WebConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config WebConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
