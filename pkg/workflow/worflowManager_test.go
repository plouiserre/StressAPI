package workflow

import (
	"testing"

	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	result "github.com/plouiserre/stressapi/pkg/result"
)

func TestCallCallApi(t *testing.T) {
	/*req_one := configuration.Request{
		Id:         1,
		Verb:       "GET",
		Uri:        "UriFirst",
		Parameters: []string{"1", "2"},
		Body:       "",
		Times:      5,
	}
	req_two := configuration.Request{
		Id:         2,
		Verb:       "PUT",
		Uri:        "UriSecond",
		Parameters: []string{"666"},
		Body:       "Body",
		Times:      10,
	}
	wf_one := configuration.Workflow{
		Id:       1,
		Requests: []configuration.Request{req_one, req_two},
	}
	req_third := configuration.Request{
		Id:         3,
		Verb:       "POST",
		Uri:        "UriThird",
		Parameters: []string{"zefazreg", "zergazeg", "aoizef"},
		Body:       "ozaehf",
		Times:      1,
	}
	req_fourth := configuration.Request{
		Id:         2,
		Verb:       "DELETE",
		Uri:        "UriFourth",
		Parameters: []string{"666"},
		Body:       "Body",
		Times:      1,
	}
	wf_second := configuration.Workflow{
		Id:       2,
		Requests: []configuration.Request{req_third, req_fourth},
	}
	conf := configuration.Configuration{
		StoreFolder: "storefolder",
		Workflows:   []configuration.Workflow{wf_one, wf_second},
	}*/

	conf := GenerateConfiguration()

	calledTimes := 4

	wm := WorkflowManager{
		Configuration: conf,
	}
	api := http.ManageApiMock{}
	resultMock := result.ResultManagerMock{}
	wm.HandleWorkflows(&api, &resultMock)

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
	conf := GenerateConfiguration()
	wm := WorkflowManager{}

	wm.SetConfigurations(conf)

	if len(wm.Configuration.Workflows) != 2 {
		t.Fatalf("length of Confs from workflowManager is 2 and not %d", len(wm.Configuration.Workflows))
	}
	if wm.Configuration.StoreFolder != "storefolder" {
		t.Fatalf("StoreFolder value is storefolder and no %s", wm.Configuration.StoreFolder)
	}
}

func GenerateConfiguration() configuration.Configuration {
	req_one := configuration.Request{
		Id:         1,
		Verb:       "GET",
		Uri:        "UriFirst",
		Parameters: []string{"1", "2"},
		Body:       "",
		Times:      5,
	}
	req_two := configuration.Request{
		Id:         2,
		Verb:       "PUT",
		Uri:        "UriSecond",
		Parameters: []string{"666"},
		Body:       "Body",
		Times:      10,
	}
	wf_one := configuration.Workflow{
		Id:       1,
		Requests: []configuration.Request{req_one, req_two},
	}
	req_third := configuration.Request{
		Id:         3,
		Verb:       "POST",
		Uri:        "UriThird",
		Parameters: []string{"zefazreg", "zergazeg", "aoizef"},
		Body:       "ozaehf",
		Times:      1,
	}
	req_fourth := configuration.Request{
		Id:         2,
		Verb:       "DELETE",
		Uri:        "UriFourth",
		Parameters: []string{"666"},
		Body:       "Body",
		Times:      1,
	}
	wf_second := configuration.Workflow{
		Id:       2,
		Requests: []configuration.Request{req_third, req_fourth},
	}
	conf := configuration.Configuration{
		StoreFolder: "storefolder",
		Workflows:   []configuration.Workflow{wf_one, wf_second},
	}
	return conf
}
