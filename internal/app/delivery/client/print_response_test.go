package client

import (
	"bytes"
	"fmt"
	"mail-test-task/internal/app/mocks"
	"mail-test-task/internal/app/models"
	"testing"
)

func TestClient_PrintResponse_Ok(t *testing.T) {
	connection := &mocks.Conn{}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	data := models.TestResponseClientOk(t)
	if err := client.PrintResponse(buf, data); err != nil {
		t.Fatal("error testing on PrintResponse with ok body")
	}

	expected := fmt.Sprintf("client_id: %s\nclient_type: %d\nexpires_in: %d\nuser_id: %d\nusername: %s\n",
		data.ClientId, data.ClientType, data.ExpiresIn, data.UserId, data.UserName)
	if buf.String() != expected {
		t.Fatalf("invalid test PrintResponsewith ok body \nexpected: %s\nreceive: %s", expected, buf.String())
	}
}
func TestClient_PrintResponse_Error(t *testing.T) {
	connection := &mocks.Conn{}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	data := models.TestResponseClientError(t)
	if err := client.PrintResponse(buf, data); err != nil {
		t.Fatal("error testing on PrintResponse with error body")
	}

	expected := fmt.Sprintf("error: %s\nmessage: %s\n", errorCodes[data.ReturnCode], data.ErrorString)
	if buf.String() != expected {
		t.Fatalf("invalid test PrintResponse with error body\nexpected: %s\nreceive: %s", expected, buf.String())
	}
}
func TestClient_PrintResponse_ErrorWithUnknownError(t *testing.T) {
	connection := &mocks.Conn{}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	data := models.TestResponseClientError(t)
	data.ReturnCode = 123
	if err := client.PrintResponse(buf, data); err != nil {
		t.Fatal("error testing on PrintResponse with unknown error")
	}
	errorString := "unknown error"
	expected := fmt.Sprintf("error: %s\nmessage: %s\n", errorString, data.ErrorString)
	if buf.String() != expected {
		t.Fatalf("invalid test PrintResponse with unknown error body\nexpected: %s\nreceive: %s",
			expected, buf.String())
	}
}
func TestClient_PrintResponse_ErrorWithUnknownRespond(t *testing.T) {
	connection := &mocks.Conn{}
	client := NewClient(connection)
	var output string
	buf := bytes.NewBufferString(output)
	data := models.TestResponseError(t)
	if err := client.PrintResponse(buf, data); err != nil {
		t.Fatal("error testing on PrintResponse with unknown error")
	}
	errorString := "response error"
	expected := fmt.Sprintf("%s", errorString)
	if buf.String() != expected {
		t.Fatalf("invalid test PrintResponse with unknown error body\nexpected: %s\nreceive: %s",
			expected, buf.String())
	}
}
