package client

import connection2 "mail-test-task/internal/app/connection"

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

	client.PrintResponse(response)
	return nil
}
