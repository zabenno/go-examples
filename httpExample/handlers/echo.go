package handlers

import (
	"log/slog"
	"net/http"
	"os"
	"io"
	"errors"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))


func EchoRequest(w http.ResponseWriter, r *http.Request) {
	responseBytes, err := echo(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Write(responseBytes)
}

func echo(r *http.Request) ([]byte, error) {
	responeBytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if len(responeBytes) == 0 {
		return nil, errors.New("Nothing given to echo")
	}

	return responeBytes, nil
}