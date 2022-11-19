package configuration

import (
	"io"
	"io/ioutil"
)

type Configurationhelper struct {
}

func (ch Configurationhelper) ReadAllIoutil(body io.Reader) ([]byte, error) {
	responseData, errData := ioutil.ReadAll(body)
	return responseData, errData
}
