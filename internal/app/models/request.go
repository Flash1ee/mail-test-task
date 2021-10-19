package models

type Request struct {
	Header Header
	SvcMsg int32
	Token  string
	Scope  string
}
