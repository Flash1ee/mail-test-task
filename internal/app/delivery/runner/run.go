package runner

import (
	"fmt"
	"mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/delivery/client"
	"os"
)

func Run(host string, port string, token string, scope string) int {
	conn, err := connection.NewTcpConnection(host, port)
	if err != nil {
		fmt.Println(err.Error())
		return CONNECTION_ERROR
	}

	cli := client.NewClient(conn)

	if err = cli.Send(token, scope); err != nil {
		fmt.Println(err.Error())
		return SEND_ERROR
	}

	if err = cli.GetResponse(os.Stdout); err != nil {
		fmt.Println(err.Error())
		return RESPONSE_ERROR
	}

	return OK
}
