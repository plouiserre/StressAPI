package main

import (
	conf "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	result "github.com/plouiserre/stressapi/pkg/result"
	wf "github.com/plouiserre/stressapi/pkg/workflow"
)

func main() {
	jsonFile := conf.JsonFile{}
	jsonFile.GetConfigurationsFromJson("../../configuration.json")
	confJson := jsonFile.GetConfigurations()
	workflowManager := wf.WorkflowManager{}
	
	//TODO tout mettre dans une méthode de workflowmanager
	workflowManager.Confs = make([]conf.Configuration, len(confJson))
	for i, conf := range  confJson{
		workflowManager.Confs[i] = *conf
	}
	
	api := http.ManageApi{}
	response := result.ResultManager{		
		StoreFolder: confJson[0].StoreFolder,
	}
	workflowManager.HandleRequests(&api, &response)
}
