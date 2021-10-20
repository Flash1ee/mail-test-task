package client

import (
	connection2 "mail-test-task/internal/app/connection"
	"os"
)

func Request(host string, port string, token string, scope string) error {
	connection, err := connection2.NewTcpConnection(host, port)
	if err != nil {
		return err
	}
	client := NewClient(connection)
	if err != nil {
		return err
	}

	response, err := client.Send(token, scope)
	if response != nil {
		return err
	}

	if err := client.PrintResponse(os.Stdout, response); err != nil {
		return err
	}

	return nil
}
