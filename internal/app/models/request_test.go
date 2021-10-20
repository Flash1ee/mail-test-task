package models

import "testing"

func cmpProtoRequest(r Request, r2 Request) bool {
	if r.SvcMsg != r2.SvcMsg {
		return false
	}
	return cmpProtoString(r.Token, r2.Token) && cmpProtoString(r.Scope, r2.Scope)
}

func TestRequest_Encode_Decode(t *testing.T) {
	testRequest := TestRequest(t)

	testEncoded, err := testRequest.Encode()
	if err != nil {
		t.Fatal("invalid encoded request")
	}
	var res Request
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded request")
	}
	if !cmpProtoRequest(testRequest, res) {
		t.Fatalf("invalid encode-decode request\nexpected: %v received: %v",
			testRequest, res)
	}
}
func TestRequest_Encode_Decode_Empty(t *testing.T) {
	testRequest := Request{}

	testEncoded, err := testRequest.Encode()
	if err != nil {
		t.Fatal("invalid encoded request")
	}
	var res Request
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded request")
	}
	if !cmpProtoRequest(testRequest, res) {
		t.Fatalf("invalid encode-decode request\nexpected: %v received: %v",
			testRequest, res)
	}
}
