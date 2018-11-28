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
	values, err := config.Get()
	checkError(err, "Could Not Parse Port & Other Configuration")
	fmt.Println(values.Port)

	// 2. register routes

	// 3. start server
}
