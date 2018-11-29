package handlers

import "fmt"
import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Sweet Home")
}
