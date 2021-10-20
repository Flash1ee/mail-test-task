package models

import "testing"

func cmpProtoResponseClientOk(r Response, r2 ResponseClientOk) bool {
	if r.ReturnCode != r2.ReturnCode {
		return false
	}
	unpackR, ok := r.Body.(*ResponseOk)
	if !ok {
		return false
	}
	return unpackR.ClientType == r2.ClientType && unpackR.ExpiresIn == r2.ExpiresIn && unpackR.UserId == r2.UserId &&
		cmpProtoString(ConvertToProtoString(r2.ClientId), unpackR.ClientId) && cmpProtoString(ConvertToProtoString(r2.UserName), unpackR.UserName)
}
func cmpProtoResponseClientError(r Response, r2 ResponseClientError) bool {
	if r.ReturnCode != r2.ReturnCode {
		return false
	}
	unpackR, ok := r.Body.(*ResponseError)
	if !ok {
		return false
	}
	return cmpProtoString(ConvertToProtoString(r2.ErrorString), unpackR.ErrorString)
}

func TestConvertToClientResponse_Invalid_ErrCode(t *testing.T) {
	response := TestResponseWithCodeOk(t)
	response.ReturnCode = -1
	if _, err := ConvertToClientResponse(response); err == nil {
		t.Fatalf("not correct convert response with negative returnCode ")
	}
}
func TestConvertToClientResponse_Ok(t *testing.T) {
	response := TestResponseWithCodeOk(t)
	res, err := ConvertToClientResponse(response)
	if err != nil {
		t.Fatal("invalid convert client response")
	}
	if !cmpProtoResponseClientOk(response, res.(ResponseClientOk)) {
		t.Fatalf("invalid encode-decode clientResponse\nexpected: %v received: %v",
			response, res)
	}
}
func TestConvertToClientResponse_Ok_Nil_Body(t *testing.T) {
	response := TestResponseWithCodeOk(t)
	response.Body = nil
	_, err := ConvertToClientResponse(response)
	if err == nil {
		t.Fatal("invalid convert client response")
	}
}
func TestConvertToClientResponse_Error(t *testing.T) {
	response := TestResponseWithError(t)
	res, err := ConvertToClientResponse(response)
	if err != nil {
		t.Fatal("invalid convert client response")
	}
	if !cmpProtoResponseClientError(response, res.(ResponseClientError)) {
		t.Fatalf("invalid encode-decode clientResponse\nexpected: %v received: %v",
			response, res)
	}
}
func TestConvertToClientResponse_Error_Nil_Body(t *testing.T) {
	response := TestResponseWithError(t)
	response.Body = nil
	_, err := ConvertToClientResponse(response)
	if err == nil {
		t.Fatal("invalid convert client response")
	}
}
