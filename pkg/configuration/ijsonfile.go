package configuration

type IJsonfile interface {
	GetConfigurationFromsJson(nameFile string)
	GetConfigurations() []*Configuration
}
