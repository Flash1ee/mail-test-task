package models

type Protocol interface {
	Decode()
	Encode()
}
