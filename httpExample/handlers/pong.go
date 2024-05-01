package handlers

import (
	"net/http"
)

func PingPongRequest(w http.ResponseWriter, r *http.Request) {
	var response []byte = []byte("Pong")
	w.Write(response)
}
