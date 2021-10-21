package main

import (
	"errors"
	"flag"
	"fmt"
	"mail-test-task/internal/app/delivery/runner"
	"os"
)

const FLAG_ERROR = -1

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 4 {
		err := errors.New("invalid args, usage ./main host port token scope")
		fmt.Println(err.Error())
		os.Exit(FLAG_ERROR)
	}

	host, port, token, scope := args[0], args[1], args[2], args[3]

	os.Exit(runner.Run(host, port, token, scope))
}
