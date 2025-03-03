package handlers

import (
	"net/http"
)

// HandleRequest is an example handler function
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
