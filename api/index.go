package handler

import (
	"fmt"
	"net/http"
)

var message = "HELLO"

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	fmt.Fprintf(w, "{'message': '%s'", message)
}
