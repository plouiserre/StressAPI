package main

import (
	"fmt"

	conf "github.com/plouiserre/stressapi/pkg/configuration"
	http "github.com/plouiserre/stressapi/pkg/http"
)

func main() {
	api := http.ManageApi{}
	confFile := conf.Configurationhelper{}
	helper := http.Httphelper{}
	jsonFile := conf.JsonFile{}
	result := api.CallApi(&jsonFile, helper, &confFile)
	fmt.Println("Response Api")
	fmt.Println(result)
}
