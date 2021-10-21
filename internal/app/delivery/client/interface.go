package client

import "io"

type IClient interface {
	Send(token string, scope string) error
	GetResponse(writer io.Writer) error
}
