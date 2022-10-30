package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type helper struct {
}

func (hp helper) GetHttp(uri string) (*http.Response, error) {
	response, err := http.Get(uri)

	return response, err
}

func (hp helper) PostHttp(uri string, json_data []byte) (*http.Response, error) {
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))

	return resp, err
}

func (hp helper) NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, uri, bytes.NewBuffer(json_data))

	return req, err
}

func (hp helper) ReadAllIoutil(body io.Reader) ([]byte, error) {
	responseData, errData := ioutil.ReadAll(body)
	return responseData, errData
}

func (hp helper) DoClient(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
