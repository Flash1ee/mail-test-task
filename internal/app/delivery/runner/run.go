package runner

import (
	"fmt"
	"mail-test-task/internal/app/connection"
	"mail-test-task/internal/app/connection/tcp_connection"
	"mail-test-task/internal/app/delivery/client"
	"os"
)

func Run(host string, port string, token string, scope string) int {
	conn, err := tcp_connection.NewTcpConnection(host, port)
	if err != nil {
		fmt.Println(err.Error())
		return CONNECTION_RESOLVE_ERROR
	}

	openConn, err := conn.Dial()
	if err != nil {
		fmt.Println(err.Error())
		return CONNECTION_OPEN_ERROR
	}

	defer func(openConn connection.Connection) {
		err := openConn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(openConn)

	cli := client.NewClient(openConn)

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
