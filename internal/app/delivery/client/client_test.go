package client

import (
	"errors"
	"mail-test-task/internal/app/mocks"
	"mail-test-task/internal/app/models"
	"testing"
)

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
