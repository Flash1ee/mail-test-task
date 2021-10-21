package client

import (
	"fmt"
	"io"
	"mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/models"
	"mail-test-task/internal/app/packet"
)

type Client struct {
	conn connection.Connection
}

func NewClient(conn connection.Connection) IClient {
	return &Client{
		conn: conn,
	}
}
func (c *Client) Send(token string, scope string) error {
	conn, err := c.conn.Dial()
	if err != nil {
		return err
	}

	defer func(conn connection.Connection) {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	if err = c.sendPackage(token, scope); err != nil {
		return err
	}
	return nil
}

func (c *Client) sendPackage(token string, scope string) error {
	data, err := packet.GetPacket(token, scope)
	if err != nil {
		return err
	}
	if _, err = c.conn.Write(data); err != nil {
		return err
	}
	return nil
}
func (c *Client) GetResponse(writer io.Writer) error {
	response, err := c.getResponse()
	if err != nil {
		return err
	}
	switch v := response.(type) {
	case models.ResponseClientOk:
		return printOk(writer, v)
	case models.ResponseClientError:
		return printErr(writer, v)
	default:
		if _, err = fmt.Fprintf(writer, "response error"); err != nil {
			return err
		}
	}
	return nil
}
func (c *Client) getResponse() (interface{}, error) {
	data, err := c.conn.Read()
	if err != nil {
		return nil, err
	}
	var res models.Response
	if err = res.Decode(data); err != nil {
		return nil, err
	}

	resp, err := models.ConvertToClientResponse(res)
	if err != nil {
		return nil, err
	}
	return resp, err
}
