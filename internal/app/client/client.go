package client

import (
	"mail-test-task/internal/app/connection"
)

type Client struct {
	conn connection.Conn
}

func (c *Client) Send(token string, scope string) (interface{}, error) {
	return nil, nil
}

func NewClient(host string, port string) (IClient, error) {
	conn, err := connection.NewTcpConnection(host, port)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}
