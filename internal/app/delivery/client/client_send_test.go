package client

import (
	conn "mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/mocks"
	"testing"
)

func TestClient_Send_InvalidWrite(t *testing.T) {
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(true),
		MockClose: mocks.MockClose(nil),
	}

	dial := mocks.MockDial(connection, nil)
	connection.MockDial = dial

	client := NewClient(connection)

	if err := client.Send("vk", "mail.ru"); err != conn.WriteError {
		t.Fatalf("error %v happened", err)
	}
}

func TestClient_Send_Success(t *testing.T) {
	connection := &mocks.Conn{
		MockWrite: mocks.MockWrite(false),
		MockClose: mocks.MockClose(nil),
	}

	dial := mocks.MockDial(connection, nil)
	connection.MockDial = dial

	client := NewClient(connection)

	if err := client.Send("vk", "mail.ru"); err != nil {
		t.Fatalf("error %v happened", err)
	}
}
