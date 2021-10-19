package main

import (
	"errors"
	"flag"
	"fmt"
	"mail-test-task/internal/app/client"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 4 {
		err := errors.New("invalid args, usage ./main host port token scope")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	host, port, token, scope := args[0], args[1], args[2], args[3]
	retCode := 0
	err := client.Run(host, port, token, scope)
	if err != nil {
		retCode = 1
	}
	os.Exit(retCode)
}
