package tcp_connection

import (
	"io/ioutil"
	"mail-test-task/internal/app/connection"
	"net"
)

type TcpConnection struct {
	tcpAddr *net.TCPAddr
	conn    *net.TCPConn
}

func NewTcpConnection(host string, port string) (connection.Connection, error) {
	addr := host + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, connection.BadResolve
	}

	return &TcpConnection{tcpAddr: tcpAddr}, nil
}

func (c *TcpConnection) Read() ([]byte, error) {
	res, err := ioutil.ReadAll(c.conn)
	if err != nil {
		return nil, connection.ReadError
	}

	return res, nil
}
func (c *TcpConnection) Close() error {
	if err := c.conn.Close(); err != nil {
		return connection.CloseError
	}

	return nil
}

func (c *TcpConnection) Write(data []byte) (int, error) {
	res, err := c.conn.Write(data)
	if err != nil {
		return -1, connection.WriteError
	}

	return res, err
}
func (c *TcpConnection) Dial() (connection.Connection, error) {
	conn, err := net.DialTCP("tcp", nil, c.tcpAddr)
	if err != nil {
		return nil, connection.DialError
	}

	return &TcpConnection{conn: conn}, nil
}
