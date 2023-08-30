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
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}

// * Create config instance
func New(filepath string) (*Config, error) {
	var config Config

	// * Load config file
	configFile, err := os.Open(filepath)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	// * Decode Json
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
	sb.WriteString("Server/Port: " + c.Server.Port)
	return sb.String()
}
