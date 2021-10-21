package client

import (
	"fmt"
	"io"
	"mail-test-task/internal/app/models"
)

func printOk(writer io.Writer, resp models.ResponseClientOk) error {
	res := fmt.Sprintf("client_id: %s\nclient_type: %d\nexpires_in: %d\nuser_id: %d\nusername: %s\n",
		resp.ClientId, resp.ClientType, resp.ExpiresIn, resp.UserId, resp.UserName)

	if _, err := fmt.Fprintf(writer, "%s", res); err != nil {
		return err
	}

	return nil
}

func printErr(writer io.Writer, resp models.ResponseClientError) error {
	var errorString string
	if resp.ReturnCode < 0 || int(resp.ReturnCode) > len(errorCodes) {
		errorString = "unknown error"
	} else {
		errorString = errorCodes[int(resp.ReturnCode)].Error()
	}

	if _, err := fmt.Fprintf(writer, "error: %s\nmessage: %s\n", errorString, resp.ErrorString); err != nil {
		return err
	}

	return nil
}
