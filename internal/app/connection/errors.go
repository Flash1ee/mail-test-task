package connection

import "errors"

var (
	BadResolve = errors.New("can not resolve tcp address")
	ReadError  = errors.New("can not read data from connection")
	WriteError = errors.New("can not write to connection")
	DialError  = errors.New("can not connect to address")
	CloseError = errors.New("error close connection")
)
