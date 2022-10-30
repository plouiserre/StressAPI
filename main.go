package main

import "fmt"

func main() {
	api := manageApi{}
	confFile := jsonFile{}
	helper := helper{}
	result := api.CallApi(&confFile, helper)
	fmt.Println("Response Api")
	fmt.Println(result)
}
