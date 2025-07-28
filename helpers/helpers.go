package helpers

import (
	"fmt"
	"net/http"
)

type HandlerMap = map[string]http.HandlerFunc

func HandleByMethod(methodHandlers HandlerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := methodHandlers[r.Method]; ok {
			handler(w, r)
		} else {
			w.Header().Set("Allow", getAllowedMethods(methodHandlers))
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getAllowedMethods(methodHandlers HandlerMap) string {
	methods := make([]string, 0, len(methodHandlers))
	for method := range methodHandlers {
		methods = append(methods, method)
	}
	return fmt.Sprintf("%s", methods)
}
