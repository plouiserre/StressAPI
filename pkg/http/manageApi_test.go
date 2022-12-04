package http

import (
	"testing"

	mock "github.com/plouiserre/stressapi/mock"
	confi "github.com/plouiserre/stressapi/pkg/configuration"
)

type manageApiTest struct {
	configurationMock mock.ConfigurationMock
	helperMock        mock.HttpHelperMock
	manageApi         ManageApi
}

func TestGetUriWithParameters(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "GET")
	uriWanted := "http://localhost:10000/congressmans/2/mandates"
	if manageApiTest.manageApi.Uri != uriWanted {
		t.Fatalf("Uri searched is %s and now the program return %s", uriWanted, manageApiTest.manageApi.Uri)
	}
}

func TestGetUriWithoutParameters(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(false, "GET")
	uriWanted := "http://localhost:10000/congressmans/"
	if manageApiTest.manageApi.Uri != uriWanted {
		t.Fatalf("Uri searched is %s and now the program return %s", uriWanted, manageApiTest.manageApi.Uri)
	}
}

func TestGetCongressmans(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "GET")
	resultWanted := `{"congressman":"bob"}`
	if manageApiTest.helperMock.IsGetHttpCalled == false {
		t.Fatalf("The method GetHttpCalled from HttpHelper is not called")
	}
	if manageApiTest.configurationMock.IsReadAllIoutil == false {
		t.Fatalf("The method ReadAllIoutil from HttpHelper is not called")
	}
	if manageApiTest.manageApi.responseRequest != resultWanted {
		t.Fatalf("Result returned by api is %s and now the program is returning %s", resultWanted, manageApiTest.manageApi.responseRequest)
	}
	if manageApiTest.manageApi.httpCode != 200 {
		t.Fatalf("Result returned by api is 200 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

func TestPostCongressman(t *testing.T) {
	manageApiTest := ManageApiInitialiedCallApi(true, "POST")

	if manageApiTest.helperMock.IsPostHttpCalled == false {
		t.Fatalf("The method GetHttpCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 201 {
		t.Fatalf("Result returned by api is 201 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

func TestPutCongressman(t *testing.T) {
	ManageNewRequestMethodTest(t, "PUT")
}

func TestDeleteCongressman(t *testing.T) {
	ManageNewRequestMethodTest(t, "DELETE")
}

func ManageNewRequestMethodTest(t *testing.T, verb string){
	manageApiTest := ManageApiInitialiedCallApi(true, verb)

	if manageApiTest.helperMock.IsNewRequestCalled == false {
		t.Fatalf("The method NewRequestCalled from HttpHelper is not called")
	}
	if manageApiTest.helperMock.IsDoClientCalled == false {
		t.Fatalf("The method DoClientCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 200 {
		t.Fatalf("Result returned by api is 200 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
}

func ManageApiInitialiedCallApi(isParameters bool, verb string) manageApiTest {
	api := ManageApi{}
	configurationMock := mock.ConfigurationMock{}
	helper := mock.HttpHelperMock{}
	conf := confi.Configuration{}
	conf.Uri = "http://localhost:10000/congressmans/"
	if isParameters{
		conf.Parameters = []string{"2", "mandates"}
	} else {
		conf.Parameters = []string{}
	}
	conf.Verb = verb
	conf.Body = "{\"congressman\":\"bob\"}"
	api.CallApi(conf, &helper, &configurationMock)
	manageApiTest := manageApiTest{
		manageApi:         api,
		configurationMock: configurationMock,
		helperMock:        helper,
	}
	return manageApiTest
}
