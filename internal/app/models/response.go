package models

type ResponseBody interface {
	Encode()
	Decode()
}

type Response struct {
	ReturnCode int32
	Body       ResponseBody
}

type ResponseOk struct {
	ClientId   string
	ClientType int32
	UserName   string
	ExpiresIn  int32
	UserId     int64
}
type ResponseError struct {
	ErrorString string
}
