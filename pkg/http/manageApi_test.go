// TODO restruct this files
package http

import (
	"testing"
)

type manageApiTest struct {
	confFileMock confFileMock
	helperMock   helperMock
	manageApi    manageApi
}

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
