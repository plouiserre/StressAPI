package mock

import (
	"io"
)

type ConfigurationMock struct {
	IsReadAllIoutil bool
}

func (cf *ConfigurationMock) ReadAllIoutil(body io.Reader) ([]byte, error) {
	cf.IsReadAllIoutil = true
	result := []byte(`{"congressman":"bob"}`)
	return result, nil
}
