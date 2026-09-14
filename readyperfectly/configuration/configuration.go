package configuration

import (
	"encoding/json"
	"flag"
	"os"
	"sync"
)

const defaultApplicationConfigurationFilePath string = "./appconfig.json"

var (
	commandLineArguments               ApplicationCommandLineArguments
	commandLineArgumentsInitialization sync.Once
)

func LoadApplicationConfigurationFromDefaultPath() (config ApplicationConfiguration, err error) {
	commandLineArguments = LoadCommandLineArguments()
	return LoadApplicationConfigurationFromFile(commandLineArguments.ApplicationConfigurationFilePath)
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

func LoadCommandLineArguments() ApplicationCommandLineArguments {
	commandLineArgumentsInitialization.Do(func() {
		flag.StringVar(
			&commandLineArguments.ApplicationConfigurationFilePath,
			"config",
			defaultApplicationConfigurationFilePath,
			"Path to appconfig.json file")

		flag.Parse()
	})

	return commandLineArguments
}
