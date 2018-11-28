package config

import (
	"flag"
	"log"
	"strings"
)

func GetPortFromCli() string {
	port := flag.String("port", "", "IP address")
	flag.Parse()

	//User is expected to give :8080 like input, if they give 8080
	//we'll append the required ':'
	if !strings.HasPrefix(*port, ":") {
		*port = ":" + *port
		log.Println("port is " + *port)
	}

	return *port
}
