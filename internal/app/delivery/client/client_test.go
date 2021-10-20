package client

import (
	"errors"
	"mail-test-task/internal/app/mocks"
	"mail-test-task/internal/app/models"
	"testing"
)

func cmpClientResponseOk(first models.ResponseClientOk, second models.ResponseClientOk) bool {
	return first.ReturnCode == second.ReturnCode && first.UserId == second.UserId && first.ExpiresIn == second.ExpiresIn &&
		first.UserName == second.UserName && first.ClientType == second.ClientType &&
		first.ClientId == second.ClientId
}
func TestClient_InvalidWrite(t *testing.T) {
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(true),
		MockClose: mocks.MockClose(nil),
	}
	dial := mocks.MockDial(connection, nil)
	connection.MockDial = dial
	client := NewClient(connection)

	if _, err := client.Send("vk", "mail.ru"); err == nil {
		t.Fatalf("error %v happened", err)
	}
}
func TestClient_InvalidDial(t *testing.T) {
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
	}

	dial := mocks.MockDial(connection, errors.New("error dial"))
	connection.MockDial = dial
	client := NewClient(connection)

	if _, err := client.Send("vk", "mail.ru"); err == nil {
		t.Fatalf("error %v happened", err)
	}
}
func TestClient_InvalidRead(t *testing.T) {
	readFunc := func(data []byte) ([]byte, error) {
		return nil, errors.New("error read")
	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}

	dial := mocks.MockDial(connection, nil)
	connection.MockDial = dial
	client := NewClient(connection)

	if _, err := client.Send("vk", "mail.ru"); err == nil {
		t.Fatalf("error %v happened", err)
	}
}
func TestClient_InvalidClose(t *testing.T) {
	readFunc := func(data []byte) ([]byte, error) {
		res := models.TestResponseWithCodeOk(t)
		return res.Encode()

	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(errors.New("happened error in connection close")),
		MockRead:  readFunc,
	}

	dial := mocks.MockDial(connection, nil)
	connection.MockDial = dial
	client := NewClient(connection)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	if _, err := client.Send("vk", "mail.ru"); err == nil {
		t.Fatalf("error %v happened", err)
	}
}
func TestClient_ResponseOk(t *testing.T) {
	ret := models.ResponseClientOk{
		ReturnCode: 0,
		ClientId:   "this is id",
		ClientType: 100,
		UserName:   "vk@mail.ru",
		ExpiresIn:  10 * 3600,
		UserId:     1000,
	}

	readFunc := func(data []byte) ([]byte, error) {
		req := models.Response{
			ReturnCode: ret.ReturnCode,
			Body: &models.ResponseOk{
				ClientId:   models.ConvertToProtoString("this is id"),
				ClientType: 100,
				UserName:   models.ConvertToProtoString("vk@mail.ru"),
				ExpiresIn:  10 * 3600,
				UserId:     1000,
			},
		}
		return req.Encode()
	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)
	res, err := client.Send("token", "scope")
	if err != nil {
		t.Fatalf("error %v happened", err)
	}
	data, ok := res.(models.ResponseClientOk)
	if !ok {
		t.Fatalf("invalid response data, expected ResponseClientOk")
	}
	if !cmpClientResponseOk(data, ret) {
		t.Fatalf("invalid return code")
	}
}
func TestClient_Response_ReturnCode_Error(t *testing.T) {
	ret := models.ResponseClientOk{
		ReturnCode: -1,
	}

	readFunc := func(data []byte) ([]byte, error) {
		req := models.Response{
			ReturnCode: ret.ReturnCode,
			Body:       &models.ResponseOk{},
		}
		return req.Encode()
	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)
	res, err := client.Send("token", "scope")
	if err == nil || res != nil {
		t.Fatalf("error %v happened", err)
	}
	if err != models.InvalidErrCode {
		t.Fatalf("error %v expected %v", err, models.InvalidErrCode)
	}
}

func TestClient_Response_EmptyBody(t *testing.T) {
	ret := models.ResponseClientOk{
		ReturnCode: 0,
	}

	readFunc := func(data []byte) ([]byte, error) {
		req := models.Response{
			ReturnCode: ret.ReturnCode,
			Body:       nil,
		}
		return req.Encode()
	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)
	res, err := client.Send("token", "scope")
	if err == nil || res != nil {
		t.Fatalf("error %v happened", err)
	}
	if err != models.EmptyBodyErr {
		t.Fatalf("error %v expected %v", err, models.InvalidErrCode)
	}
}
func TestClient_ResponseErrorBody(t *testing.T) {
	ret := models.ResponseClientError{
		ReturnCode:  1,
		ErrorString: "error happened",
	}

	readFunc := func(data []byte) ([]byte, error) {
		req := models.Response{
			ReturnCode: ret.ReturnCode,
			Body: &models.ResponseError{
				ErrorString: models.ConvertToProtoString(ret.ErrorString),
			},
		}
		return req.Encode()
	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)
	res, err := client.Send("token", "scope")
	if err != nil || res == nil {
		t.Fatalf("error %v happened", err)
	}

	data, ok := res.(models.ResponseClientError)
	if !ok {
		t.Fatalf("invalid response data, expected ResponseClientOk")
	}
	if data.ReturnCode != ret.ReturnCode || ret.ErrorString != data.ErrorString {
		t.Fatalf("invalid return code")
	}
}
func TestClient_Response_CheckPackingData(t *testing.T) {
	token := "token"
	scope := "scope"

	readFunc := func(data []byte) ([]byte, error) {
		var res models.Request
		err := res.Decode(data)
		if err != nil {
			return nil, err
		}

		tokenFromRead, err := res.Token.ToString()
		if err != nil || tokenFromRead != token {
			return nil, err
		}
		scopeFromRead, err := res.Scope.ToString()
		if err != nil || scopeFromRead != scope {
			return nil, err
		}
		response := models.TestResponseWithCodeOk(t)
		return response.Encode()
	}
	connection := &mocks.Conn{
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)
	res, err := client.Send(token, scope)
	if err != nil || res == nil {
		t.Fatalf("error %v happened", err)
	}
}
