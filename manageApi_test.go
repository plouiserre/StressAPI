// TODO restruct this files
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

type manageApiTest struct {
	confFileMock confFileMock
	helperMock   helperMock
	manageApi    manageApi
}

// TODO externalise in confilemock.go
type confFileMock struct {
	configuration *Configuration
	isParameters  bool
	Verb          string
}

func (cf confFileMock) GetConfigurationFromJson(nameFile string) {
	fmt.Print("lol")
}

func (cf confFileMock) GetConfiguration() *Configuration {
	cf.configuration = &Configuration{}
	cf.configuration.Uri = "http://localhost:10000/congressmans/"
	if cf.isParameters == true {
		cf.configuration.Parameters = []string{"2", "mandates"}
	} else {
		cf.configuration.Parameters = []string{}
	}

	cf.configuration.Verb = cf.Verb

	cf.configuration.Body = "{\"congressman\":\"bob\"}"

	return cf.configuration
}

// TODO externalise in helpermock.go
// -----------------Début helpermock -----------------//
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

//-----------------Fin helpermock -----------------//

func TestGetUriWithParameters(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "GET")
	uriWanted := "http://localhost:10000/congressmans/2/mandates"
	if manageApiTest.manageApi.uri != uriWanted {
		t.Fatalf("Uri searched is %s and now the program return %s", uriWanted, manageApiTest.manageApi.uri)
	}
}

func TestGetUriWithoutParameters(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(false, "GET")
	uriWanted := "http://localhost:10000/congressmans/"
	if manageApiTest.manageApi.uri != uriWanted {
		t.Fatalf("Uri searched is %s and now the program return %s", uriWanted, manageApiTest.manageApi.uri)
	}
}

func TestGetCongressmans(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "GET")
	resultWanted := `{"congressman":"bob"}`
	if manageApiTest.helperMock.isGetHttpCalled == false {
		t.Fatalf("The method GetHttpCalled from HttpHelper is not called")
	}
	if manageApiTest.helperMock.isReadAllIoutil == false {
		t.Fatalf("The method ReadAllIoutil from HttpHelper is not called")
	}
	if manageApiTest.manageApi.result != resultWanted {
		t.Fatalf("Result returned by api is %s and now the program is returning %s", resultWanted, manageApiTest.manageApi.result)
	}
	if manageApiTest.manageApi.httpCode != 200 {
		t.Fatalf("Result returned by api is 200 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

func TestPostCongressman(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "POST")

	if manageApiTest.helperMock.isPostHttpCalled == false {
		t.Fatalf("The method GetHttpCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 201 {
		t.Fatalf("Result returned by api is 201 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

// TODO factoriser TestPutCongressman et TestDeleteCongressman
func TestPutCongressman(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "PUT")

	if manageApiTest.helperMock.isNewRequestCalled == false {
		t.Fatalf("The method NewRequestCalled from HttpHelper is not called")
	}
	if manageApiTest.helperMock.isDoClientCalled == false {
		t.Fatalf("The method DoClientCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 200 {
		t.Fatalf("Result returned by api is 200 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

func TestDeleteCongressman(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "DELETE")

	if manageApiTest.helperMock.isNewRequestCalled == false {
		t.Fatalf("The method NewRequestCalled from HttpHelper is not called")
	}
	if manageApiTest.helperMock.isDoClientCalled == false {
		t.Fatalf("The method DoClientCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 200 {
		t.Fatalf("Result returned by api is 200 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

func ManageApiInitialiedCallApi(isParameters bool, verb string) manageApiTest {
	api := manageApi{}
	confFile := confFileMock{}
	confFile.isParameters = isParameters
	confFile.Verb = verb
	helper := helperMock{}
	api.CallApi(confFile, &helper)
	manageApiTest := manageApiTest{
		manageApi:    api,
		confFileMock: confFile,
		helperMock:   helper,
	}
	return manageApiTest
}
