package models

import "testing"

func cmpProtoResponseError(r ResponseError, r2 ResponseError) bool {
	return cmpProtoString(r.ErrorString, r2.ErrorString)
}
func cmpProtoResponseOk(r ResponseOk, r2 ResponseOk) bool {
	if !cmpProtoString(r.ClientId, r2.ClientId) || !cmpProtoString(r.UserName, r2.UserName) {
		return false
	}

	return r.UserId == r2.UserId && r.ClientType == r2.ClientType && r.ExpiresIn == r2.ExpiresIn
}
func cmpProtoResponse(r Response, r2 Response) bool {
	if r.ReturnCode != r2.ReturnCode {
		return false
	}

	if r.ReturnCode == 0 {
		return cmpProtoResponseOk(*r.Body.(*ResponseOk), *r2.Body.(*ResponseOk))
	} else {
		return cmpProtoResponseError(*r.Body.(*ResponseError), *r2.Body.(*ResponseError))
	}
}
func TestResponseError_Encode_Decode(t *testing.T) {
	expResponseError := TestResponseError(t)

	testEncoded, err := expResponseError.Encode()
	if err != nil {
		t.Fatal("invalid encoded responseError")
	}

	var res ResponseError
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded responseError")
	}

	if !cmpProtoResponseError(expResponseError, res) {
		t.Fatalf("invalid encode-decode responseError\nexpected: %v received: %v",
			expResponseError, res)
	}
}

func TestResponseError_Encode_Decode_Empty(t *testing.T) {
	expResponseError := ResponseError{}

	testEncoded, err := expResponseError.Encode()
	if err != nil {
		t.Fatal("invalid encoded responseError")
	}

	var res ResponseError
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded responseError")
	}

	if !cmpProtoResponseError(expResponseError, res) {
		t.Fatalf("invalid encode-decode responseError\nexpected: %v received: %v",
			expResponseError, res)
	}
}
func TestResponseOk_Encode_Decode(t *testing.T) {
	expResponseOk := TestResponseOk(t)

	testEncoded, err := expResponseOk.Encode()
	if err != nil {
		t.Fatal("invalid encoded responseOk")
	}

	var res ResponseOk
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded responseOk")
	}

	if !cmpProtoResponseOk(expResponseOk, res) {
		t.Fatalf("invalid encode-decode responseError\nexpected: %v received: %v",
			expResponseOk, res)
	}
}
func TestResponseOk_Encode_Decode_Empty(t *testing.T) {
	expResponseOk := ResponseOk{}
	testEncoded, err := expResponseOk.Encode()

	if err != nil {
		t.Fatal("invalid encoded responseOk")
	}

	var res ResponseOk
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded responseOk")
	}

	if !cmpProtoResponseOk(expResponseOk, res) {
		t.Fatalf("invalid encode-decode responseError\nexpected: %v received: %v",
			expResponseOk, res)
	}
}
func TestResponse_Encode_Decode_Ok(t *testing.T) {
	expResponseOk := TestResponseWithCodeOk(t)

	testEncoded, err := expResponseOk.Encode()
	if err != nil {
		t.Fatal("invalid encoded responseOk")
	}

	var res Response
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded responseOk")
	}

	if !cmpProtoResponse(expResponseOk, res) {
		t.Fatalf("invalid encode-decode response\nexpected: %v received: %v",
			expResponseOk, res)
	}
}
func TestResponse_Encode_Decode_Ok_Empty(t *testing.T) {
	expResponseOk := Response{}
	expResponseOk.ReturnCode = 0x00000000

	_, err := expResponseOk.Encode()
	if err == nil {
		t.Fatal("invalid encoded responseOk")
	}
}
func TestResponse_Encode_Decode_Err(t *testing.T) {
	expResponseOk := TestResponseWithError(t)

	testEncoded, err := expResponseOk.Encode()
	if err != nil {
		t.Fatal("invalid encoded responseOk")
	}

	var res Response
	if err = res.Decode(testEncoded); err != nil {
		t.Fatal("invalid decoded responseOk")
	}

	if !cmpProtoResponse(expResponseOk, res) {
		t.Fatalf("invalid encode-decode response\nexpected: %v received: %v",
			expResponseOk, res)
	}
}
func TestResponse_Encode_Decode_Err_Empty(t *testing.T) {
	expResponseOk := Response{}
	expResponseOk.ReturnCode = 0x00000001

	_, err := expResponseOk.Encode()
	if err == nil {
		t.Fatal("invalid encoded responseOk")
	}
}
