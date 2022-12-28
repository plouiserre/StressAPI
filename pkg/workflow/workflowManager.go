package workflow

import (
	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	resultPkg "github.com/plouiserre/stressapi/pkg/result"
)

// https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1
type WorkflowManager struct {
	Configuration configuration.Configuration
	api           http.IManageApi
	response      resultPkg.IResultManager
}

func (wf *WorkflowManager) SetConfigurations(configuration configuration.Configuration) {
	wf.Configuration = configuration
}

func (wf WorkflowManager) HandleWorkflows(api http.IManageApi, response resultPkg.IResultManager) {
	wf.api = api
	wf.response = response
	for _, workflow := range wf.Configuration.Workflows {
		for _, request := range workflow.Requests {
			wf.HandleRequest(request)
		}
	}
}

func (wf WorkflowManager) HandleRequest(req configuration.Request) {
	confFile := configuration.Configurationhelper{}
	helper := http.Httphelper{}
	results := wf.api.CallApis(req, helper, &confFile)
	for _, result := range results {
		wf.response.SetResult(result)
		wf.response.StoreResult()
	}
}
