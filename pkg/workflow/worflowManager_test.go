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

func testHandlesRequest(t *testing.T, calledTimes int){
	wm := WorkflowManager{
		Confs : []configuration.Configuration{},	
	}
	wm.Confs = make([]configuration.Configuration, calledTimes)
	
	api := http.ManageApiMock{}
	resultMock := result.ResultManagerMock{}
	wm.HandleRequests(&api, &resultMock)	
	
	if api.IsCallApiCalling == false{
		t.Fatalf("Method CallApi is not call in the test TestCallCallApi")
	}
	
	if api.NumberCalled != calledTimes{
		t.Fatalf("Method CallApis must be called only %d time(s) and not %d", calledTimes, api.NumberCalled)
	}
	
	if resultMock.IsStoreResultCalled == false{
		t.Fatalf("Method StoreResult is not call in the test TestCallCallApi")
	}
}