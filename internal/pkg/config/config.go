package config

import (
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	Logger struct {
		LogLevel string `json:"log-level"`
	} `json:"logger"`
}

func New(filepath string) (*Config, error) {
	var config Config

	// * Load config file
	configFile, err := os.Open(filepath)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) String() string {
	sb := strings.Builder{}
	sb.WriteString("Logger/LogLevel: " + c.Logger.LogLevel)
	return sb.String()
}
