package configuration

type IConfigurationHelper interface {
	GetConfigurationFromJson(nameFile string)
	GetConfiguration() *Configuration
}
