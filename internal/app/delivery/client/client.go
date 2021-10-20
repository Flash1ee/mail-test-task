package client

import (
	"fmt"
	"io"
	"mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/models"
)

type Client struct {
	conn connection.Conn
}

func NewClient(conn connection.Conn) *Client {
	return &Client{
		conn: conn,
	}
}
func (c *Client) Send(token string, scope string) (interface{}, error) {
	conn, err := c.conn.Dial()
	if err != nil {
		return nil, err
	}

	defer func(conn connection.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

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
	return resp, err
}
func (c *Client) PrintResponse(writer io.Writer, resp interface{}) error {
	switch v := resp.(type) {
	case models.ResponseClientOk:
		return printOk(writer, v)
	case models.ResponseClientError:
		return printErr(writer, v)
	default:
		if _, err := fmt.Fprintf(writer, "response error"); err != nil {
			return err
		}
	}
	return nil
}
