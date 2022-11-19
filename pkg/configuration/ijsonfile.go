package configuration

type IJsonfile interface {
	GetConfigurationFromJson(nameFile string)
	GetConfiguration() *Configuration
}
