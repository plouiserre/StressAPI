// TODO tester le nombre d'appel avec times
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
	apiCalledTimes := 1
	manageApiTest := ManageApiInitialiedCallApi(true, "GET", apiCalledTimes)
	uriWanted := "http://localhost:10000/congressmans/2/mandates"
	if manageApiTest.manageApi.Uri != uriWanted {
		t.Fatalf("Uri searched is %s and now the program return %s", uriWanted, manageApiTest.manageApi.Uri)
	}
	if manageApiTest.helperMock.CallTimes != apiCalledTimes{
		t.Fatalf("Called Api is called %d and we want  %d called", manageApiTest.helperMock.CallTimes, apiCalledTimes)
	}
}

func TestGetUriWithoutParameters(t *testing.T) {
	apiCalledTimes := 2
	manageApiTest := ManageApiInitialiedCallApi(false, "GET", apiCalledTimes)
	uriWanted := "http://localhost:10000/congressmans/"
	if manageApiTest.manageApi.Uri != uriWanted {
		t.Fatalf("Uri searched is %s and now the program return %s", uriWanted, manageApiTest.manageApi.Uri)
	}
	if manageApiTest.helperMock.CallTimes != apiCalledTimes{
		t.Fatalf("Called Api is called %d and we want  %d called", manageApiTest.helperMock.CallTimes, apiCalledTimes)
	}
}

func TestGetCongressmans(t *testing.T) {
	apiCalledTimes := 3
	manageApiTest := ManageApiInitialiedCallApi(true, "GET", apiCalledTimes)
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
	if manageApiTest.helperMock.CallTimes != apiCalledTimes{
		t.Fatalf("Called Api is called %d and we want  %d called", manageApiTest.helperMock.CallTimes, apiCalledTimes)
	}
}

func TestPostCongressman(t *testing.T) {
	apiCalledTimes := 4
	manageApiTest := ManageApiInitialiedCallApi(true, "POST", apiCalledTimes)

	if manageApiTest.helperMock.IsPostHttpCalled == false {
		t.Fatalf("The method GetHttpCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 201 {
		t.Fatalf("Result returned by api is 201 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
	if manageApiTest.helperMock.CallTimes != apiCalledTimes{
		t.Fatalf("Called Api is called %d and we want  %d called", manageApiTest.helperMock.CallTimes, apiCalledTimes)
	}
}

func TestPutCongressman(t *testing.T) {
	ManageNewRequestMethodTest(t, "PUT")
}

func TestDeleteCongressman(t *testing.T) {
	ManageNewRequestMethodTest(t, "DELETE")
}

func ManageNewRequestMethodTest(t *testing.T, verb string){
	apiCalledTimes := 5
	manageApiTest := ManageApiInitialiedCallApi(true, verb, apiCalledTimes)

	if manageApiTest.helperMock.IsNewRequestCalled == false {
		t.Fatalf("The method NewRequestCalled from HttpHelper is not called")
	}
	if manageApiTest.helperMock.IsDoClientCalled == false {
		t.Fatalf("The method DoClientCalled from HttpHelper is not called")
	}

	if manageApiTest.manageApi.httpCode != 200 {
		t.Fatalf("Result returned by api is 200 and now the program is returning %d", manageApiTest.manageApi.httpCode)
	}
	if manageApiTest.helperMock.CallTimes != apiCalledTimes{
		t.Fatalf("Called Api is called %d and we want  %d called", manageApiTest.helperMock.CallTimes, apiCalledTimes)
	}
}

func ManageApiInitialiedCallApi(isParameters bool, verb string, timesCalled int) manageApiTest {
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
	conf.Times = timesCalled
	api.CallApis(conf, &helper, &configurationMock)
	manageApiTest := manageApiTest{
		manageApi:         api,
		configurationMock: configurationMock,
		helperMock:        helper,
	}
	return manageApiTest
}
//TODO finir cette page 
