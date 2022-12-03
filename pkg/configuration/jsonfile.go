package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonFile struct {
	configuration *Configuration
}

func (jf *JsonFile) GetConfigurationFromJson(nameFile string) {
	var conf Configuration

	file, err := os.Open(nameFile)

	if err != nil {
		fmt.Println(err.Error())
	}

	data, _ := ioutil.ReadAll(file)

	json.Unmarshal(data, &conf)

	jf.configuration = &conf

	defer file.Close()
}

func (jf *JsonFile) GetConfiguration() *Configuration {
	return jf.configuration
}
