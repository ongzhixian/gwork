package configuration

import (
	"encoding/json"
	"os"
)

const defaultApplicationConfigurationFilePath string = "./appconfig.json"

func LoadApplicationConfigurationFromDefaultPath() (config ApplicationConfiguration, err error) {
	return LoadApplicationConfigurationFromFile(defaultApplicationConfigurationFilePath)
}

func LoadApplicationConfigurationFromFile(filePath string) (config ApplicationConfiguration, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	return config, err
}
