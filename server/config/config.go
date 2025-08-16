package config

import (
	"encoding/json"
	"os"
)

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	var configJSON string
	if env != "PROD" {
		data, _ := os.ReadFile("config.example.json")
		configJSON = string(data)
	} else {
		configJSON = os.Getenv("CONFIG_JSON")
	}
	var config Config
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		return nil, err
	}
	return &config, nil
}
