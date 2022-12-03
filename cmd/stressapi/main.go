package main

import (
	"fmt"

	conf "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
	resultPkg "github.com/plouiserre/stressapi/pkg/result"
)

func main() {
	api := http.ManageApi{}
	confFile := conf.Configurationhelper{}
	helper := http.Httphelper{}
	jsonFile := conf.JsonFile{}
	jsonFile.GetConfigurationFromJson("../../configuration.json")
	conf := *jsonFile.GetConfiguration()
	result := api.CallApi(conf, helper, &confFile)
	fmt.Println("Response Api")
	fmt.Println(result)
	response := resultPkg.ResultManager{		
		Result : result,
		StoreFolder: conf.StoreFolder,
	}
	response.StoreResult()
}
