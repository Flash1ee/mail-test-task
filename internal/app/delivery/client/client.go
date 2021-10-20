package client

import (
	"fmt"
	"mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/models"
)

type Client struct {
	conn connection.Conn
}

func NewClient(host string, port string) (IClient, error) {
	conn, err := connection.NewTcpConnection(host, port)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}
func (c *Client) Send(token string, scope string) (interface{}, error) {
	conn, err := c.conn.Dial()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	if err := sendPackage(conn, token, scope); err != nil {
		return nil, err
	}

	respond, err := getRespond(conn)
	if err != nil {
		return nil, err
	}

	return respond, nil
}

func sendPackage(conn connection.Conn, token string, scope string) error {
	data, err := getPackage(token, scope)
	if err != nil {
		return err
	}
	if _, err := conn.Write(data); err != nil {
		return err
	}
	return nil
}
func getRespond(conn connection.Conn) (interface{}, error) {
	data, err := conn.Read()
	if err != nil {
		return nil, err
	}
	var res models.Response
	if err := res.Decode(data); err != nil {
		return nil, err
	}

	resp, err := models.ConvertToClientResponse(res)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (c *Client) PrintResponse(resp interface{}) {
	switch v := resp.(type) {
	case models.ResponseClientOk:
		printOk(v)
	case models.ResponseClientError:
		printErr(v)
	default:
		fmt.Println("response error")
	}
}
func printOk(resp models.ResponseClientOk) {
	fmt.Printf("client_id: %s\n", resp.ClientId)
	fmt.Printf("clint_type: %d\n", resp.ClientType)
	fmt.Printf("expires_in: %d\n", resp.ExpiresIn)
	fmt.Printf("user_id: %d\n", resp.UserId)
	fmt.Printf("username: %s\n", resp.UserName)
}

func printErr(resp models.ResponseClientError) {
	var errorString string
	if resp.ReturnCode < 0 || int(resp.ReturnCode) > len(errors) {
		errorString = "unknown error"
	} else {
		errorString = errors[resp.ReturnCode]
	}
	fmt.Printf("error: %s\n", errorString)
	fmt.Printf("message: %s\n", resp.ErrorString)
}
