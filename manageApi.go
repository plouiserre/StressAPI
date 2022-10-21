package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type manageApi struct {
	//result *string
}

func (ma manageApi) CallApi() string {
	response, err := http.Get("http://localhost:10000/congressmans/")

	if err != nil {
		fmt.Println(err)
	}

	responseData, errData := ioutil.ReadAll(response.Body)

	if errData != nil {
		fmt.Println(errData)
	}

	result := (string(responseData))

	//ma.result = &result

	return result
}
