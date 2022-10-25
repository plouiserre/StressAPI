package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type manageApi struct {
	configuration Configuration
	result        string
}

func (ma *manageApi) CallApi() string {
	confFile := jsonFile{}

	confFile.GetConfiguration("configuration.json")

	ma.configuration = *confFile.configuration

	if ma.configuration.Verb == "GET" {
		ma.CallGetEndpoint()
	} else if ma.configuration.Verb == "POST" {
		ma.CallPostEndpoint()
	}

	return ma.result
}

func (ma *manageApi) CallGetEndpoint() {
	uri := ma.GetCompleteUri()
	response, err := http.Get(uri)

	if err != nil {
		fmt.Println(err)
	}

	responseData, errData := ioutil.ReadAll(response.Body)

	if errData != nil {
		fmt.Println(errData)
	}

	ma.result = (string(responseData))

	fmt.Println(response.StatusCode)
}

func (ma *manageApi) GetCompleteUri() string {
	uri := ma.configuration.Uri
	for i := 0; i < len(ma.configuration.Parameters); i++ {
		uri += "/" + ma.configuration.Parameters[i]
	}
	return uri
}

func (ma *manageApi) CallPostEndpoint() {
	var body map[string]string
	err_json := json.Unmarshal([]byte(ma.configuration.Body), &body)
	if err_json != nil {
		log.Fatal(err_json)
	} else {
		json_data, err_marshal := json.Marshal(body)
		if err_marshal != nil {
			log.Fatal(err_marshal)
		} else {
			resp, err := http.Post(ma.configuration.Uri, "application/json", bytes.NewBuffer(json_data))

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(resp.StatusCode)
		}
	}
}
