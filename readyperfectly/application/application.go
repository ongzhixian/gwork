package application

import (
	"readyperfectly/configuration"
)

func DoWork() string {
	return "Do some work"
}

func GetApplicationConfiguration() (config configuration.ApplicationConfiguration, err error) {
	return configuration.LoadApplicationConfigurationFromDefaultPath()
}
