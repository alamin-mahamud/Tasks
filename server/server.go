package server

import (
	"log"
	"net/http"
)

func Start(port string) {
	log.Println("running server on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
