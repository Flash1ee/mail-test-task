package models

import (
	"testing"
)

func cmpProtoHeader(h Header, h2 Header) bool {
	return h.SvcId == h2.SvcId && h.BodyLen == h2.BodyLen && h.RequestId == h2.RequestId
}
func TestHeader_Encode_Decode(t *testing.T) {
	expectedHeader := TestHeader(t)
	testHeader := expectedHeader

	binHeader, err := testHeader.Encode()
	if err != nil {
		t.Fatal("invalid header encode")
	}

	var res Header
	if err = res.Decode(binHeader); err != nil {
		t.Fatal("invalid header decode")
	}

	if !cmpProtoHeader(expectedHeader, res) {
		t.Fatalf("invalid encode-decode header\nexpected: %v received: %v",
			expectedHeader, res)
	}
}
func TestHeader_Encode_Decode_Empty(t *testing.T) {
	expectedHeader := Header{}
	testHeader := expectedHeader

	binHeader, err := testHeader.Encode()
	if err != nil {
		t.Fatal("invalid header encode")
	}

	var res Header
	if err = res.Decode(binHeader); err != nil {
		t.Fatal("invalid header decode")
	}

	if !cmpProtoHeader(expectedHeader, res) {
		t.Fatalf("invalid encode-decode header\nexpected: %v received: %v",
			expectedHeader, res)
	}
}
