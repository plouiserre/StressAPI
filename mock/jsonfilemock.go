package mock

import (
	"fmt"

	conf "github.com/plouiserre/stressapi/pkg/configuration"
)

type JsonFileMock struct {
	configuration *conf.Request
	IsParameters  bool
	Verb          string
}

func (jf JsonFileMock) GetConfigurationFromJson(nameFile string) {
	fmt.Print("lol")
}

func (jf JsonFileMock) GetConfiguration() *conf.Request {
	jf.configuration = &conf.Request{}
	jf.configuration.Uri = "http://localhost:10000/congressmans/"
	if jf.IsParameters {
		jf.configuration.Parameters = []string{"2", "mandates"}
	} else {
		jf.configuration.Parameters = []string{}
	}

	jf.configuration.Verb = jf.Verb

	jf.configuration.Body = "{\"congressman\":\"bob\"}"

	return jf.configuration
}
