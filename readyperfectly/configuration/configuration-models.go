package configuration

type ApplicationConfiguration struct {
	AwsConfiguration AwsConfiguration `json:"aws"`
}

type AwsConfiguration struct {
	ProfileName string `json:"profileName"`
}

type DBConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
}

type ApplicationCommandLineArguments struct {
	ApplicationConfigurationFilePath string
}
