package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pbzona/vercel-go/helpers"
)

type ApiResponse struct {
	message string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	helpers.HandleByMethod(helpers.HandlerMap{
		"GET": handleGet,
	})
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	data := ApiResponse{
		message: "Hi from Go running on Vercel",
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		http.Error(w, "Failed to encode API response", 500)
	}

	buf.WriteTo(w)
}
