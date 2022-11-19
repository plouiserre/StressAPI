package http

import "net/http"

type IHttpHelper interface {
	GetHttp(uri string) (*http.Response, error)
	PostHttp(uri string, json_data []byte) (*http.Response, error)
	NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error)
	DoClient(req *http.Request) (*http.Response, error)
}
