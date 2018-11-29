package routes

import (
	"net/http"

	"github.com/alamin-mahamud/go-tasks/handlers"
)

func Register() {
	http.HandleFunc("/", handlers.Home)
}
