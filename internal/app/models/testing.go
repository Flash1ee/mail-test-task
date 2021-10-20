package models

import "testing"

func TestToken(t *testing.T) String {
	t.Helper()

	token := "token"
	return GetIprotoString(token)
}
func TestScope(t *testing.T) String {
	t.Helper()

	scope := "scope"
	return GetIprotoString(scope)
}
func TestRequest(t *testing.T) Request {
	t.Helper()
	return Request{
		Header: TestHeader(t),
		SvcMsg: 10,
		Token:  TestToken(t),
		Scope:  TestScope(t),
	}
}
func TestHeader(t *testing.T) Header {
	t.Helper()

	return Header{
		SvcId:     1,
		BodyLen:   2,
		RequestId: 3,
	}
}
func TestResponseError(t *testing.T) ResponseError {
	t.Helper()
	errorMsg := "happened error"
	return ResponseError{
		ErrorString: GetIprotoString(errorMsg),
	}
}
func TestResponseOk(t *testing.T) ResponseOk {
	t.Helper()
	clientId := "1"
	username := "username"

	return ResponseOk{
		ClientId:   GetIprotoString(clientId),
		ClientType: 1,
		UserName:   GetIprotoString(username),
		ExpiresIn:  10,
		UserId:     2,
	}
}
func TestResponseWithCodeOk(t *testing.T) Response {
	t.Helper()
	responseOk := TestResponseOk(t)
	return Response{
		ReturnCode: 0x00000000,
		Body:       &responseOk,
	}
}
func TestResponseWithError(t *testing.T) Response {
	t.Helper()
	responseError := TestResponseError(t)
	return Response{
		ReturnCode: 0x00000001,
		Body:       &responseError,
	}
}
