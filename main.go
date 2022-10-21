package main

import "fmt"

func main() {
	api := manageApi{}
	result := api.CallApi()
	fmt.Println("Response Api")
	fmt.Println(result)
}
