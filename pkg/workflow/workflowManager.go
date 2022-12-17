package workflow

import (
	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	resultPkg "github.com/plouiserre/stressapi/pkg/result"
)

type WorkflowManager struct {
	Confs []configuration.Configuration	
	api http.IManageApi
	response resultPkg.IResultManager
}

func (wf WorkflowManager) HandleRequests(api http.IManageApi, response resultPkg.IResultManager) {
	wf.api = api
	wf.response = response
	wf.HandleRequest()
}

func(wf WorkflowManager) HandleRequest(){
	confFile := configuration.Configurationhelper{}
	helper := http.Httphelper{}
	results := wf.api.CallApis(wf.Confs[0], helper, &confFile)
	for _, result := range results {
		wf.response.SetResult(result)
		wf.response.StoreResult()
	}
}