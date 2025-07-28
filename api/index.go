package api

import (
	"fmt"
	"net/http"

	"github.com/pbzona/vercel-go/helpers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	helpers.HandleByMethod(helpers.HandlerMap{
		"GET": handleGet,
	})
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v", r)
}
