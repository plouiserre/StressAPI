package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonFile struct {
	configuration *Configuration
}

func (jf *jsonFile) GetConfiguration(nameFile string) {
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
