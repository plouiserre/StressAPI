package main

import (
	conf "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	result "github.com/plouiserre/stressapi/pkg/result"
	wf "github.com/plouiserre/stressapi/pkg/workflow"
)

func main() {
	jsonFile := conf.JsonFile{}
	jsonFile.GetConfigurationFromJson("../../configuration.json")
	confJson := *jsonFile.GetConfiguration()
	workflowManager := wf.WorkflowManager{
		Conf : confJson,	
	}
	api := http.ManageApi{}
	response := result.ResultManager{		
		StoreFolder: confJson.StoreFolder,
	}
	workflowManager.HandleRequests(&api, &response)
}
