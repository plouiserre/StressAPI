//TODO identify actions to solve problems if I put this file in mock folder

package http

import (
	conf "github.com/plouiserre/stressapi/pkg/configuration"
	result "github.com/plouiserre/stressapi/pkg/result"
)

type ManageApiMock struct {
	IsCallApiCalling bool
}

func (mm *ManageApiMock)CallApi(configuration conf.Configuration, httpHelper IHttpHelper, confHelper conf.IConfigurationHelper) []result.Result{		
	results  := make([]result.Result, 1)
	resultMock := result.Result{
		HttpCode: 200,
	}
	results[0] = resultMock
	
	mm.IsCallApiCalling = true
	return results
}
func (mm ManageApiMock)CallGetEndpoint(){
	
} 
func (mm ManageApiMock)CallDeleteEndpoint() {
	
}
func (mm ManageApiMock)CallPutEndpoint(){
	
}
func (mm ManageApiMock)ManageNewRequest(httpMethod string, json_data []byte){
	
}
func (mm ManageApiMock)GetCompleteUri(){
	
} 
func (mm ManageApiMock)CallPostEndpoint() {
	
}
func (mm ManageApiMock)GetJsonData() (bool, []byte) {
	fake := []byte{}
	return false, fake
}
