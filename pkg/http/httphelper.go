package http

import (
	"bytes"
	"net/http"
)

type Httphelper struct {
}

func (hp Httphelper) GetHttp(uri string) (*http.Response, error) {
	response, err := http.Get(uri)

	return response, err
}

func (hp Httphelper) PostHttp(uri string, json_data []byte) (*http.Response, error) {
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))

	return resp, err
}

func (hp Httphelper) NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, uri, bytes.NewBuffer(json_data))

	return req, err
}

func (hp Httphelper) DoClient(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
