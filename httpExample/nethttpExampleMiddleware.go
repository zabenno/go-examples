package main

import (
	"httpexample/handlers"
	requestlog "httpexample/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

func main() {
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	log.Printf("Starting net/http example server")
	httpServer := http.NewServeMux()

	echoHandler := http.HandlerFunc(handlers.EchoRequest)
	versionEndpointHandler := http.HandlerFunc(handlers.VersionEndpoint)
	pingpongHandler := http.HandlerFunc(handlers.PingPongRequest)

	httpServer.Handle("GET /version", requestlog.Middleware(versionEndpointHandler))
	httpServer.Handle("POST /echo", requestlog.Middleware(echoHandler))
	httpServer.Handle("GET /ping", requestlog.Middleware(pingpongHandler))

	if err := http.ListenAndServe("localhost:8080", httpServer); err != nil {
		logger.Error(err.Error())
	}
}
