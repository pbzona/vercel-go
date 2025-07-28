package handler

import (
	"fmt"
	"net/http"
)

var message = "This is the testing endpoint"

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	fmt.Fprintf(w, "{'message': '%s'", message)
}

