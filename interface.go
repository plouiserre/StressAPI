package main

import (
	"io"
	"net/http"
)

type IConfiguration interface {
	GetConfigurationFromJson(nameFile string)
	GetConfiguration() *Configuration
}

type IHelper interface {
	GetHttp(uri string) (*http.Response, error)
	PostHttp(uri string, json_data []byte) (*http.Response, error)
	NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error)
	ReadAllIoutil(body io.Reader) ([]byte, error)
	DoClient(req *http.Request) (*http.Response, error)
}
