package packet

import "errors"

var (
	InvalidPackingHeader = errors.New("invalid packing body")
	InvalidPackingBody   = errors.New("invalid packing body")
)
