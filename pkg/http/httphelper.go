package http

import (
	"bytes"
	"net/http"
)

type httphelper struct {
}

func (hp httphelper) GetHttp(uri string) (*http.Response, error) {
	response, err := http.Get(uri)

	return response, err
}

func (hp httphelper) PostHttp(uri string, json_data []byte) (*http.Response, error) {
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))

	return resp, err
}

func (hp httphelper) NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, uri, bytes.NewBuffer(json_data))

	return req, err
}

func (hp httphelper) DoClient(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
