package workflow

import (
	"testing"

	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	result "github.com/plouiserre/stressapi/pkg/result"
)

func TestHandleOneRequest(t *testing.T) {
	testHandlesRequest(t, 1)
}

func TestHandleMoreThanOneRequest(t *testing.T) {
	testHandlesRequest(t, 3)
}

func testHandlesRequest(t *testing.T, calledTimes int) {
	wm := WorkflowManager{
		Confs: []configuration.Configuration{},
	}
	wm.Confs = make([]configuration.Configuration, calledTimes)

	api := http.ManageApiMock{}
	resultMock := result.ResultManagerMock{}
	wm.HandleRequests(&api, &resultMock)

	if api.IsCallApiCalling == false {
		t.Fatalf("Method CallApi is not call in the test TestCallCallApi")
	}

	if api.NumberCalled != calledTimes {
		t.Fatalf("Method CallApis must be called only %d time(s) and not %d", calledTimes, api.NumberCalled)
	}

	if resultMock.IsStoreResultCalled == false {
		t.Fatalf("Method StoreResult is not call in the test TestCallCallApi")
	}
}

func SetConfigurations(t *testing.T) {
	firstConf := configuration.Configuration{
		Verb:        "GET",
		Uri:         "http://localhost:10000/congressmans/",
		Parameters:  []string{},
		Body:        "",
		StoreFolder: "/Users/plouiserre/Desktop/Save",
		Times:       4,
	}
	secondConf := configuration.Configuration{
		Verb:        "POST",
		Uri:         "http://localhost:10000/congressmans/",
		Parameters:  []string{},
		Body:        "{\"Uid\": \"PA666440\",\"Civility\": \"M.\",\"FirstName\": \"Pierre-Louis\",\"LastName\": \"Serré\",\"Alpha\": \"SERRE66\",\"Trigramme\": \"PLS\",\"BirthDate\": \"1989-03-28 00:00:00\",\"BirthCity\": \"Versailles\",\"BirthDepartment\": \"Yvelines\",\"BirthCountry\": \"France\",\"Jobtitle\": \"Ingénieur\",\"CatSocPro\": \"Professions geek\",\"FamSocPro\": \"Cadres et professions intellectuelles supérieures\"}",
		StoreFolder: "/Users/plouiserre/Desktop/Save",
		Times:       1,
	}
	confs := []configuration.Configuration{
		firstConf,
		secondConf,
	}
	wm := WorkflowManager{}

	wm.SetConfigurations(confs)

	if len(wm.Confs) != 2 {
		t.Fatalf("length of Confs from workflowManager is 2 and not %d", len(wm.Confs))
	}
	if wm.Confs[0].Body != "GET" || wm.Confs[1].Body != "POST" {
		t.Fatalf("Set of Confs are broken")
	}
}
