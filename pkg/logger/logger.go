package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Logger wraps slog.Logger with convenience methods.
type Logger struct {
	*slog.Logger
}

// New creates a new Logger at the given level ("debug", "info", "warn", "error").
func New(level string) *Logger {
	var lvl slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl})
	return &Logger{slog.New(handler)}
}

// Infof logs a formatted info message.
func (l *Logger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

// Fatalf logs a formatted error message then exits.
func (l *Logger) Fatalf(format string, args ...any) {
	l.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}
