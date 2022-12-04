package workflow

import (
	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	resultPkg "github.com/plouiserre/stressapi/pkg/result"
)

type WorkflowManager struct {
	Conf configuration.Configuration	
}

func (wf WorkflowManager) HandleRequests(api http.IManageApi, response resultPkg.IResultManager) {
	confFile := configuration.Configurationhelper{}
	helper := http.Httphelper{}
	result := api.CallApi(wf.Conf, helper, &confFile)
	response.SetResult(result)
	response.StoreResult()
}