package requestlog

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

func Middleware(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handling Request", "Path", r.URL.Path, "Method", r.Method)

		startTime := time.Now()
		lrw := newLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, r)
		endTime := time.Since(startTime)

		logger.Info("Finished Request", "Path", r.URL.Path, "Method", r.Method, "StatusCode", lrw.statusCode, "ExecutionTime", endTime.Milliseconds())
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}
