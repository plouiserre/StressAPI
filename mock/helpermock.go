package mock

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpHelperMock struct {
	IsGetHttpCalled    bool
	IsPostHttpCalled   bool
	IsNewRequestCalled bool
	IsDoClientCalled   bool
	CallTimes 		   int 
}

func (hm *HttpHelperMock) GetHttp(uri string) (*http.Response, error) {
	hm.IsGetHttpCalled = true
	hm.CallTimes += 1
	responbody := ioutil.NopCloser(bytes.NewBuffer([]byte(`{"value":"fixed"}`)))
	return &http.Response{
		StatusCode: 200,
		Body:       responbody,
	}, nil
}

func (hm *HttpHelperMock) PostHttp(uri string, json_data []byte) (*http.Response, error) {
	hm.IsPostHttpCalled = true
	hm.CallTimes += 1
	return &http.Response{
		StatusCode: 201,
		Body:       nil,
	}, nil
}

func (hm *HttpHelperMock) NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error) {
	hm.IsNewRequestCalled = true
	hm.CallTimes += 1
	return &http.Request{}, nil
}

func (hm *HttpHelperMock) DoClient(req *http.Request) (*http.Response, error) {
	hm.IsDoClientCalled = true
	responbody := ioutil.NopCloser(bytes.NewBuffer([]byte(`{"value":"fixed"}`)))
	return &http.Response{
		StatusCode: 200,
		Body:       responbody,
	}, nil
}
