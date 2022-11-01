package main

import (
	"fmt"

	_ "github.com/plouiserre/stressapi/http"
)

func main() {
	api := manageApi{}
	confFile := jsonFile{}
	helper := helper{}
	result := api.CallApi(&confFile, helper)
	fmt.Println("Response Api")
	fmt.Println(result)
}

//source pour le découpage https://medium.com/sellerapp/golang-project-structuring-ben-johnson-way-2a11035f94bc
