package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonFile struct {
	configurations []*Configuration
}

func (jf *JsonFile) GetConfigurationsFromJson(nameFile string) {
	var confs []Configuration

	file, err := os.Open(nameFile)

	if err != nil {
		fmt.Println(err.Error())
	}

	data, _ := ioutil.ReadAll(file)

	json.Unmarshal(data, &confs)

	jf.configurations = make([]*Configuration, len(confs))
	
	for i, conf := range confs{
		jf.configurations[i] = &conf
	}

	defer file.Close()
}

func (jf *JsonFile) GetConfigurations() []*Configuration {
	return jf.configurations
}
