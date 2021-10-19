package client

type IClient interface {
	Send(token string, scope string) (interface{}, error)
}
