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
