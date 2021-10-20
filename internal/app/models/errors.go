package models

import "errors"

var (
	EmptyBodyErr   = errors.New("empty body")
	InvalidErrCode = errors.New("invalid response code")
)
