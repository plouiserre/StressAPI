package main

import (
	"fmt"
	"io/ioutil"
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
}

func (ma *manageApi) GetCompleteUri() string {
	uri := ma.configuration.Uri
	for i := 0; i < len(ma.configuration.Parameters); i++ {
		uri += "/" + ma.configuration.Parameters[i]
	}
	return uri
}
