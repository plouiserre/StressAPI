package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//TODO manage when the server is not answered add in readme

type manageApi struct {
	configuration Configuration
	result        string
	helper        IHelper
	uri           string
	httpCode      int
}

func (ma *manageApi) CallApi(confFile IConfiguration, helper IHelper) string {
	ma.helper = helper

	confFile.GetConfigurationFromJson("configuration.json")

	ma.configuration = *confFile.GetConfiguration()

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
	ma.GetCompleteUri()

	response, err := ma.helper.GetHttp(ma.uri)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	ma.httpCode = response.StatusCode

	responseData, errData := ma.helper.ReadAllIoutil(response.Body)

	if errData != nil {
		fmt.Println(errData)
	}

	ma.result = (string(responseData))

	fmt.Println(response.StatusCode)
}

func (ma *manageApi) CallDeleteEndpoint() {
	ma.GetCompleteUri()
	ma.ManageNewRequest(http.MethodDelete, nil)
}

func (ma *manageApi) CallPutEndpoint() {
	ma.GetCompleteUri()
	is_ok, json_data := ma.GetJsonData()
	if is_ok {
		ma.ManageNewRequest(http.MethodPut, json_data)
	}
}

func (ma *manageApi) ManageNewRequest(httpMethod string, json_data []byte) {
	req, err_new_request := ma.helper.NewRequestHttp(httpMethod, ma.uri, json_data)

	if err_new_request != nil {
		log.Fatal(err_new_request)
	}

	resp, err := ma.helper.DoClient(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	ma.httpCode = resp.StatusCode

	fmt.Println(resp.StatusCode)
}

func (ma *manageApi) GetCompleteUri() {
	uri := ma.configuration.Uri
	for i := 0; i < len(ma.configuration.Parameters); i++ {
		if i == 0 {
			uri += ma.configuration.Parameters[i]
		} else {
			uri += "/" + ma.configuration.Parameters[i]
		}
	}
	ma.uri = uri
}

func (ma *manageApi) CallPostEndpoint() {
	is_ok, json_data := ma.GetJsonData()
	if is_ok {
		resp, err := ma.helper.PostHttp(ma.configuration.Uri, json_data)

		if err != nil {
			log.Fatal(err)
		}

		ma.httpCode = resp.StatusCode

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
