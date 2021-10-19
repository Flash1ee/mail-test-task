package connection

type Conn interface {
	Dial() (Conn, error)
	Read() ([]byte, error)
	Write(data []byte) (int, error)
	Close() error
}
