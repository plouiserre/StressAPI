package http

import (
	conf "github.com/plouiserre/stressapi/pkg/configuration"
	result "github.com/plouiserre/stressapi/pkg/result"
)

type IManageApi interface {
	CallApis(configuration conf.Request, httpHelper IHttpHelper, confHelper conf.IConfigurationHelper) []result.Result
	CallGetEndpoint()
	CallDeleteEndpoint()
	CallPutEndpoint()
	ManageNewRequest(httpMethod string, json_data []byte)
	GetCompleteUri()
	CallPostEndpoint()
	GetJsonData() (bool, []byte)
}
