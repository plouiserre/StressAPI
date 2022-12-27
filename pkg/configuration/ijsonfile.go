package configuration

type IJsonfile interface {
	GetConfigurationsFromJson(nameFile string)
	GetConfigurations() []*Configuration
}
