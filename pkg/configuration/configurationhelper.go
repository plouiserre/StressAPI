package configuration

import (
	"io"
	"io/ioutil"
)

type configurationhelper struct {
}

func (ch configurationhelper) ReadAllIoutil(body io.Reader) ([]byte, error) {
	responseData, errData := ioutil.ReadAll(body)
	return responseData, errData
}
