package connection

type Connection interface {
	Dial() (Connection, error)
	Read() ([]byte, error)
	Write(data []byte) (int, error)
	Close() error
}
