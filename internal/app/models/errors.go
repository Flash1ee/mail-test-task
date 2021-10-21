package models

import "errors"

var (
	EmptyBodyErr   = errors.New("empty body")
	InvalidErrCode = errors.New("invalid response code")
	InvalidEncode  = errors.New("invalid encoded data")
	InvalidDecode  = errors.New("invalid decoded data")
)
