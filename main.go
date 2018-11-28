package main

import (
	"fmt"

	"github.com/alamin-mahamud/go-tasks/config"
)

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}

func main() {
	// 1. get configs and override with cli flags
	values, err := config.Get("config/config.json")
	checkError(err, "Could Not Parse Port & Other Configuration")

	// get port from env if config is not present
	if err != nil {
		port := config.GetPortFromCli()
		values.Port = port
	}

	fmt.Println("Hello " + values.Port)

	// 2. register routes

	// 3. start server
}
