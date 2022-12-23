package workflow

import (
	configuration "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	resultPkg "github.com/plouiserre/stressapi/pkg/result"
)

type WorkflowManager struct {
	Confs    []configuration.Configuration
	api      http.IManageApi
	response resultPkg.IResultManager
}

// TODO faire la copie après
func (wf *WorkflowManager) SetConfigurations(confs []configuration.Configuration) {
	wf.Confs = make([]configuration.Configuration, len(confs))
	for i, conf := range confs {
		wf.Confs[i] = conf
	}
}

func (wf WorkflowManager) HandleRequests(api http.IManageApi, response resultPkg.IResultManager) {
	wf.api = api
	wf.response = response
	for _, conf := range wf.Confs {
		wf.HandleRequest(conf)
	}
}

func (wf WorkflowManager) HandleRequest(conf configuration.Configuration) {
	confFile := configuration.Configurationhelper{}
	helper := http.Httphelper{}
	results := wf.api.CallApis(conf, helper, &confFile)
	for _, result := range results {
		wf.response.SetResult(result)
		wf.response.StoreResult()
	}
}
