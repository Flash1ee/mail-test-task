package mocks

import "mail-test-task/internal/app/connection"

type Conn struct {
	MockRead  ReadFunc
	MockWrite WriteFunc
	MockDial  Dial
	MockClose CloseFunc
	data      []byte
}

func (c *Conn) Dial() (connection.Connection, error) {
	return c.MockDial()
}

func (c *Conn) Read() ([]byte, error) {
	if c.MockRead != nil {
		return c.MockRead(c.data)
	}
	return c.data, nil
}
func (c *Conn) Write(data []byte) (int, error) {
	if c.MockWrite != nil {
		return c.MockWrite(data)
	}
	c.data = data
	return len(c.data), nil
}
func (c *Conn) Close() error {
	if c.MockClose != nil {
		return c.MockClose()
	}
	return nil
}
