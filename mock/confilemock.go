package mock

import "fmt"

// TODO externalise in confilemock.go
type confFileMock struct {
	configuration *Configuration
	isParameters  bool
	Verb          string
}

func (cf confFileMock) GetConfigurationFromJson(nameFile string) {
	fmt.Print("lol")
}

func (cf confFileMock) GetConfiguration() *Configuration {
	cf.configuration = &Configuration{}
	cf.configuration.Uri = "http://localhost:10000/congressmans/"
	if cf.isParameters == true {
		cf.configuration.Parameters = []string{"2", "mandates"}
	} else {
		cf.configuration.Parameters = []string{}
	}

	cf.configuration.Verb = cf.Verb

	cf.configuration.Body = "{\"congressman\":\"bob\"}"

	return cf.configuration
}
