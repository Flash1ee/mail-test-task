package mocks

import (
	"mail-test-task/internal/app/connection"
)

type Dial func() (connection.Connection, error)
type CloseFunc func() error
type ReadFunc func([]byte) ([]byte, error)
type WriteFunc func([]byte) (int, error)

const (
	MOCK_WRITE_N_OK  = 0
	MOCK_WRITE_N_ERR = -1
)

func MockClose(err error) CloseFunc {
	return func() error {
		return err
	}
}

func MockWrite(isErr bool) WriteFunc {
	return func([]byte) (int, error) {
		if isErr {
			return MOCK_WRITE_N_ERR, connection.WriteError
		}

		return MOCK_WRITE_N_OK, nil
	}
}

func MockDial(conn connection.Connection, err error) Dial {
	return func() (connection.Connection, error) {
		return conn, err
	}
}
