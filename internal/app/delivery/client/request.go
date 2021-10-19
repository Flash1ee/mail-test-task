package client

import "fmt"

func Request(host string, port string, token string, scope string) error {
	client, err := NewClient(host, port)
	if err != nil {
		return err
	}
	respond, err := client.Send(token, scope)
	if respond != nil {
		return err
	}
	fmt.Println(respond)
	return nil
}
