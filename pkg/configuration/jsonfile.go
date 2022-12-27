package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonFile struct {
	configuration Configuration
}

func (jf *JsonFile) GetConfigurationsFromJson(nameFile string) {
	var conf Configuration

	file, err := os.Open(nameFile)

	if err != nil {
		fmt.Println(err.Error())
	}

	data, _ := ioutil.ReadAll(file)

	json.Unmarshal(data, &conf)

	//jf.workflows = make([]Workflow, len(confs))

	//TODO mettre un copy à la place
	/*for i, conf := range confs {
		jf.workflows[i] = conf
	}*/
	jf.configuration = conf

	defer file.Close()
}

func (jf *JsonFile) GetConfigurations() Configuration {
	return jf.configuration
}
