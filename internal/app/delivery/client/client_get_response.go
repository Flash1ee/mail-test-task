package client

import (
	"fmt"
	"io"
	"mail-test-task/internal/app/models"
)

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
