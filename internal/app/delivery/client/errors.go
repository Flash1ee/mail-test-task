package client

type CubeError struct {
	error string
}

func New(error string) error {
	return &CubeError{error}
}

func (e *CubeError) Error() string {
	return e.error
}

var errorCodes = map[int]error{
	0x00000000: &CubeError{error: "CUBE_OAUTH2_ERR_OK"},
	0x00000001: &CubeError{error: "CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND"},
	0x00000002: &CubeError{error: "CUBE_OAUTH2_ERR_DB_ERROR"},
	0x00000003: &CubeError{error: "CUBE_OAUTH2_ERR_UNKNOWN_MSG"},
	0x00000004: &CubeError{error: "CUBE_OAUTH2_ERR_BAD_PACKET"},
	0x00000005: &CubeError{error: "CUBE_OAUTH2_ERR_BAD_CLIENT"},
	0x00000006: &CubeError{error: "CUBE_OAUTH2_ERR_BAD_SCOPE"},
}
