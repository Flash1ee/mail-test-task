package client

import (
	"bytes"
	"fmt"
	conn "mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/mocks"
	"mail-test-task/internal/app/models"
	"testing"
)

func TestClient_GetResponse_InvalidRead(t *testing.T) {
	readFunc := func(data []byte) ([]byte, error) {
		return nil, conn.ReadError
	}
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}

	dial := mocks.MockDial(connection, nil)
	connection.MockDial = dial
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	if err := client.GetResponse(buf); err != conn.ReadError {
		t.Fatalf("error %v happened", err)
	}
}

func TestClient_GetResponse_Ok(t *testing.T) {
	res := models.TestResponseWithCodeOk(t)
	clientType := res.Body.(*models.ResponseOk).ClientType
	expiresIn := res.Body.(*models.ResponseOk).ExpiresIn
	clientId, err := res.Body.(*models.ResponseOk).ClientId.ToString()
	if err != nil {
		t.Fatal(err.Error())
	}
	userName, err := res.Body.(*models.ResponseOk).UserName.ToString()
	if err != nil {
		t.Fatal(err.Error())
	}
	userId := res.Body.(*models.ResponseOk).UserId

	expected := models.ResponseClientOk{
		ReturnCode: res.ReturnCode,
		ClientId:   clientId,
		ClientType: clientType,
		ExpiresIn:  expiresIn,
		UserName:   userName,
		UserId:     userId,
	}
	readFunc := func(data []byte) ([]byte, error) {
		return res.Encode()
	}
	connection := &mocks.Conn{
		MockRead: readFunc,
	}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	if err = client.GetResponse(buf); err != nil {
		t.Fatal("error testing on GetResponse with ok body")
	}

	expectedRes := fmt.Sprintf("client_id: %s\nclient_type: %d\nexpires_in: %d\nuser_id: %d\nusername: %s\n",
		expected.ClientId, expected.ClientType, expected.ExpiresIn, expected.UserId, expected.UserName)
	if buf.String() != expectedRes {
		t.Fatalf("invalid test PrintResponsewith ok body \nexpected: %s\nreceive: %s", expectedRes, buf.String())
	}
}
func TestClient_GetResponse_Error(t *testing.T) {
	res := models.TestResponseWithError(t)
	expErrorString, err := res.Body.(*models.ResponseError).ErrorString.ToString()
	if err != nil {
		t.Fatal(err.Error())
	}

	readFunc := func(data []byte) ([]byte, error) {
		return res.Encode()
	}
	connection := &mocks.Conn{
		MockRead: readFunc,
	}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	if err := client.GetResponse(buf); err != nil {
		t.Fatal("error testing on GetResponse with error body")
	}

	expectedRes := fmt.Sprintf("error: %s\nmessage: %s\n", errorCodes[res.ReturnCode], expErrorString)
	if buf.String() != expectedRes {
		t.Fatalf("invalid test GetResponse with error body\nexpected: %s\nreceive: %s", expectedRes, buf.String())
	}
}
func TestClient_GetResponse_ErrorWithUnknownError(t *testing.T) {
	res := models.TestResponseWithError(t)
	res.ReturnCode = 123
	expErrorString, err := res.Body.(*models.ResponseError).ErrorString.ToString()
	if err != nil {
		t.Fatal(err.Error())
	}

	readFunc := func(data []byte) ([]byte, error) {
		return res.Encode()
	}
	connection := &mocks.Conn{
		MockRead: readFunc,
	}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	if err := client.GetResponse(buf); err != nil {
		t.Fatal("error testing on GetResponse with unknown error")
	}
	errorString := "unknown error"
	expectedRes := fmt.Sprintf("error: %s\nmessage: %s\n", errorString, expErrorString)
	if buf.String() != expectedRes {
		t.Fatalf("invalid test GetResponse with unknown error body\nexpected: %s\nreceive: %s",
			expectedRes, buf.String())
	}
}

// незаявленный протоколом ответ
func TestClient_GetResponse_Error_IncorrectServerRespond(t *testing.T) {
	res := models.TestResponseError(t)
	readFunc := func(data []byte) ([]byte, error) {
		return res.Encode()
	}
	connection := &mocks.Conn{
		MockRead: readFunc,
	}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	if err := client.GetResponse(buf); err != models.InvalidDecode {
		t.Fatal("error testing on GetResponse with unknown error")
	}
}
