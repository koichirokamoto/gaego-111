package handlers

import (
	"net/http"
)

// Hello responses hello message.
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
