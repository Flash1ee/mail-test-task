package client

func Request(host string, port string, token string, scope string) error {
	client, err := NewClient(host, port)
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
