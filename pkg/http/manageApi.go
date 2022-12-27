package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	conf "github.com/plouiserre/stressapi/pkg/configuration"
	result "github.com/plouiserre/stressapi/pkg/result"
)

type ManageApi struct {
	request         conf.Request
	responseRequest string
	httpHelper      IHttpHelper
	httpCode        int
	confHelper      conf.IConfigurationHelper
	results         []result.Result
	Uri             string
}

func (ma *ManageApi) CallApis(req conf.Request, httpHelper IHttpHelper, confHelper conf.IConfigurationHelper) []result.Result {

	ma.httpHelper = httpHelper

	ma.confHelper = confHelper

	ma.request = req

	ma.results = make([]result.Result, ma.request.Times)

	for i := 0; i < ma.request.Times; i++ {
		ma.CallApi(i)
	}

	return ma.results
}

func (ma *ManageApi) CallApi(timeCalled int) {
	start := time.Now()
	if ma.request.Verb == "GET" {
		ma.CallGetEndpoint()
	} else if ma.request.Verb == "POST" {
		ma.CallPostEndpoint()
	} else if ma.request.Verb == "DELETE" {
		ma.CallDeleteEndpoint()
	} else if ma.request.Verb == "PUT" {
		ma.CallPutEndpoint()
	} else {
		fmt.Println("Error verb unknown")
	}
	end := time.Now()

	requestDuration := end.Sub(start)

	ma.results[timeCalled] = result.Result{
		Response:        ma.responseRequest,
		HttpCode:        ma.httpCode,
		Body:            ma.request.Body,
		UriCalled:       ma.request.Uri,
		RequestDuration: requestDuration.String(),
	}
}

func (ma *ManageApi) CallGetEndpoint() {
	ma.GetCompleteUri()

	response, err := ma.httpHelper.GetHttp(ma.Uri)

	if err != nil {
		fmt.Println(err)
		ma.httpCode = 408
	} else {

		defer response.Body.Close()

		ma.httpCode = response.StatusCode

		responseData, errData := ma.confHelper.ReadAllIoutil(response.Body)

		if errData != nil {
			fmt.Println(errData)
		}

		ma.responseRequest = (string(responseData))

		fmt.Println(response.StatusCode)
	}
}

func (ma *ManageApi) CallDeleteEndpoint() {
	ma.GetCompleteUri()
	ma.ManageNewRequest(http.MethodDelete, nil)
}

func (ma *ManageApi) CallPutEndpoint() {
	ma.GetCompleteUri()
	is_ok, json_data := ma.GetJsonData()
	if is_ok {
		ma.ManageNewRequest(http.MethodPut, json_data)
	}
}

func (ma *ManageApi) ManageNewRequest(httpMethod string, json_data []byte) {
	req, err_new_request := ma.httpHelper.NewRequestHttp(httpMethod, ma.Uri, json_data)

	if err_new_request != nil {
		log.Fatal(err_new_request)
	}

	resp, err := ma.httpHelper.DoClient(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	ma.httpCode = resp.StatusCode

	fmt.Println(resp.StatusCode)
}

func (ma *ManageApi) GetCompleteUri() {
	uri := ma.request.Uri
	for i := 0; i < len(ma.request.Parameters); i++ {
		if i == 0 {
			uri += ma.request.Parameters[i]
		} else {
			uri += "/" + ma.request.Parameters[i]
		}
	}
	ma.Uri = uri
}

func (ma *ManageApi) CallPostEndpoint() {
	is_ok, json_data := ma.GetJsonData()
	if is_ok {
		resp, err := ma.httpHelper.PostHttp(ma.request.Uri, json_data)

		if err != nil {
			log.Fatal(err)
		}

		ma.httpCode = resp.StatusCode

		fmt.Println(resp.StatusCode)
	}
}

func (ma *ManageApi) GetJsonData() (bool, []byte) {
	var body map[string]string

	err_json := json.Unmarshal([]byte(ma.request.Body), &body)
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
