package workflow

import (
	"testing"

	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	result "github.com/plouiserre/stressapi/pkg/result"
)

type workflowManagerTest struct {
}

func TestCallMethods(t *testing.T) {
	workflowManager := WorkflowManager{
		Conf : configuration.Configuration{},	
	}
	
	api := http.ManageApiMock{}
	resultMock := result.ResultManagerMock{}
	workflowManager.HandleRequests(&api, &resultMock)	
	
	if api.IsCallApiCalling == false{
		t.Fatalf("Method CallApi is not call in the test TestCallCallApi")
	}
	
	if resultMock.IsStoreResultCalled == false{
		t.Fatalf("Method StoreResult is not call in the test TestCallCallApi")
	}
}