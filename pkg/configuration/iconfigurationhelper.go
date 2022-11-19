package configuration

import "io"

type IConfigurationHelper interface {
	ReadAllIoutil(body io.Reader) ([]byte, error)
}
