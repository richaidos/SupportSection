package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Unable to read configuration file %s", path)
	}

	config := new(Config)

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse configuration file %s", path)
	}

	return config, nil
}
