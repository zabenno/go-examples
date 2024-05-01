package utils

import (
	"log/slog"
	"os"
	"strings"
)

func SetupLogger(LogLevel string) slog.Logger {
	ParsedLevel := strings.ToUpper(LogLevel)
	lvl := new(slog.LevelVar)
	switch ParsedLevel {
	case "DEBUG":
		lvl.Set(slog.LevelDebug)
	case "INFO":
		lvl.Set(slog.LevelInfo)
	case "WARN":
		lvl.Set(slog.LevelWarn)
	default:
		lvl.Set(slog.LevelInfo)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl}))

	return *logger
}
