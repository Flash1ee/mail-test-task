package client

import (
	"bytes"
	"fmt"
	"mail-test-task/internal/app/mocks"
	"mail-test-task/internal/app/models"
	"testing"
)

func TestClient_SendAndResponse_Ok_Response(t *testing.T) {
	token := "token"
	scope := "scope"
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
		var req models.Request
		err = req.Decode(data)
		if err != nil {
			return nil, err
		}

		tokenFromRead, err := req.Token.ToString()
		if err != nil || tokenFromRead != token {
			return nil, err
		}

		scopeFromRead, err := req.Scope.ToString()
		if err != nil || scopeFromRead != scope {
			return nil, err
		}

		return res.Encode()
	}
	connection := &mocks.Conn{
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)

	var output string
	buf := bytes.NewBufferString(output)
	if err = client.Send(token, scope); err != nil {
		t.Fatalf("error %v happened", err)
	}
	if err = client.GetResponse(buf); err != nil {
		t.Fatalf("error %v happened", err)
	}

	expectedRes := fmt.Sprintf("client_id: %s\nclient_type: %d\nexpires_in: %d\nuser_id: %d\nusername: %s\n",
		expected.ClientId, expected.ClientType, expected.ExpiresIn, expected.UserId, expected.UserName)
	if buf.String() != expectedRes {
		t.Fatalf("invalid test PrintResponsewith ok body \nexpected: %s\nreceive: %s", expectedRes, buf.String())
	}
}
func TestClient_SendAndResponse_ErrorResponse(t *testing.T) {
	token := "token"
	scope := "scope"
	res := models.TestResponseWithError(t)

	expErrorString, err := res.Body.(*models.ResponseError).ErrorString.ToString()

	if err != nil {
		t.Fatal(err.Error())
	}
	readFunc := func(data []byte) ([]byte, error) {
		var req models.Request
		err = req.Decode(data)
		if err != nil {
			return nil, err
		}

		tokenFromRead, err := req.Token.ToString()
		if err != nil || tokenFromRead != token {
			return nil, err
		}

		scopeFromRead, err := req.Scope.ToString()
		if err != nil || scopeFromRead != scope {
			return nil, err
		}

		return res.Encode()
	}
	connection := &mocks.Conn{
		MockClose: mocks.MockClose(nil),
		MockRead:  readFunc,
	}
	connection.MockDial = mocks.MockDial(connection, nil)
	client := NewClient(connection)

	var output string
	buf := bytes.NewBufferString(output)
	if err = client.Send(token, scope); err != nil {
		t.Fatalf("error %v happened", err)
	}
	if err = client.GetResponse(buf); err != nil {
		t.Fatalf("error %v happened", err)
	}

	expectedRes := fmt.Sprintf("error: %s\nmessage: %s\n", errorCodes[int(res.ReturnCode)], expErrorString)
	if buf.String() != expectedRes {
		t.Fatalf("invalid test GetResponse with error body\nexpected: %s\nreceive: %s", expectedRes, buf.String())
	}
}
