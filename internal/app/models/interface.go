package models

type Protocol interface {
	Decode() ([]byte, error)
	Encode() error
}
