package connection

import (
	"errors"
	"io/ioutil"
	"net"
)

type TcpConnection struct {
	tcpAddr *net.TCPAddr
	conn    *net.TCPConn
}

func (c *TcpConnection) Read() ([]byte, error) {
	return ioutil.ReadAll(c.conn)
}
func (c *TcpConnection) Close() error {
	return c.conn.Close()
}

func (c *TcpConnection) Write(data []byte) (int, error) {
	return c.conn.Write(data)
}
func (c *TcpConnection) Dial() (Conn, error) {
	conn, err := net.DialTCP("tcp", nil, c.tcpAddr)
	if err != nil {
		return nil, err
	}
	return &TcpConnection{conn: conn}, nil
}
func NewTcpConnection(host string, port string) (Conn, error) {
	addr := host + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, errors.New("can not resolve tcp address")
	}
	return &TcpConnection{tcpAddr: tcpAddr}, nil
}
