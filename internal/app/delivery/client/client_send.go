package client

import (
	"mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/packet"
)

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
