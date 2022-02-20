package configuration

import (
	"fmt"
)

type InitSettings struct {
	Config *Config
}

func GetInitSettings() *InitSettings {
	iSettings := new(InitSettings)
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	iSettings.Config = config

	return iSettings
}
