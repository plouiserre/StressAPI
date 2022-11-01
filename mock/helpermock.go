package mock

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type helperMock struct {
	isGetHttpCalled    bool
	isReadAllIoutil    bool
	isPostHttpCalled   bool
	isNewRequestCalled bool
	isDoClientCalled   bool
}

func (hm *helperMock) GetHttp(uri string) (*http.Response, error) {
	hm.isGetHttpCalled = true
	responbody := ioutil.NopCloser(bytes.NewBuffer([]byte(`{"value":"fixed"}`)))
	return &http.Response{
		StatusCode: 200,
		Body:       responbody,
	}, nil
}

func (hm *helperMock) PostHttp(uri string, json_data []byte) (*http.Response, error) {
	hm.isPostHttpCalled = true
	return &http.Response{
		StatusCode: 201,
		Body:       nil,
	}, nil
}

func (hm *helperMock) NewRequestHttp(httpMethod string, uri string, json_data []byte) (*http.Request, error) {
	hm.isNewRequestCalled = true
	return &http.Request{}, nil
}

func (hm *helperMock) ReadAllIoutil(body io.Reader) ([]byte, error) {
	hm.isReadAllIoutil = true
	result := []byte(`{"congressman":"bob"}`)
	return result, nil
}

func (hm *helperMock) DoClient(req *http.Request) (*http.Response, error) {
	hm.isDoClientCalled = true
	responbody := ioutil.NopCloser(bytes.NewBuffer([]byte(`{"value":"fixed"}`)))
	return &http.Response{
		StatusCode: 200,
		Body:       responbody,
	}, nil
}
