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
	} else if ma.configuration.Verb == "DELETE" {
		ma.CallDeleteEndpoint()
	} else if ma.configuration.Verb == "PUT" {
		ma.CallPutEndpoint()
	} else {
		fmt.Println("Error verb unknown")
	}

	return ma.result
}

func (ma *manageApi) CallGetEndpoint() {
	uri := ma.GetCompleteUri()
	response, err := http.Get(uri)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	responseData, errData := ioutil.ReadAll(response.Body)

	if errData != nil {
		fmt.Println(errData)
	}

	ma.result = (string(responseData))

	fmt.Println(response.StatusCode)
}

func (ma *manageApi) CallDeleteEndpoint() {
	uri := ma.GetCompleteUri()
	ma.ManageNewRequest(http.MethodDelete, uri, nil)
}

func (ma *manageApi) CallPutEndpoint() {
	uri := ma.GetCompleteUri()
	is_ok, json_data := ma.GetJsonData()
	if is_ok {
		ma.ManageNewRequest(http.MethodPut, uri, json_data)
	}
}

func (ma *manageApi) ManageNewRequest(httpMethod string, uri string, json_data []byte) {
	req, _ := http.NewRequest(httpMethod, uri, bytes.NewBuffer(json_data))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}

func (ma *manageApi) GetCompleteUri() string {
	uri := ma.configuration.Uri
	for i := 0; i < len(ma.configuration.Parameters); i++ {
		if i == 0 {
			uri += ma.configuration.Parameters[i]
		} else {
			uri += "/" + ma.configuration.Parameters[i]
		}
	}
	return uri
}

func (ma *manageApi) CallPostEndpoint() {
	is_ok, json_data := ma.GetJsonData()
	if is_ok {
		resp, err := http.Post(ma.configuration.Uri, "application/json", bytes.NewBuffer(json_data))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.StatusCode)
	}
}

func (ma *manageApi) GetJsonData() (bool, []byte) {
	var body map[string]string

	err_json := json.Unmarshal([]byte(ma.configuration.Body), &body)
	if err_json != nil {
		log.Fatal(err_json)
		return false, nil
	} else {
		json_data, err_marshal := json.Marshal(body)
		if err_marshal != nil {
			log.Fatal(err_marshal)
			return false, nil
		} else {
			return true, json_data
		}
	}
}
