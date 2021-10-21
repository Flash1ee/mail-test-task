package client

import (
	"mail-test-task/internal/app/connection"
)

type Client struct {
	conn connection.Connection
}

func NewClient(conn connection.Connection) IClient {
	return &Client{
		conn: conn,
	}
}
