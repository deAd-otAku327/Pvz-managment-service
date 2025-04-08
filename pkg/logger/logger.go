package logger

import (
	"errors"
	"io"
	"log/slog"
)

var errInvalidLogLevel = errors.New("invalid logging level provided")

var logLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func New(w io.Writer, level string) (*slog.Logger, error) {
	lvl, ok := logLevels[level]
	if !ok {
		return nil, errInvalidLogLevel
	}
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: lvl})), nil
}
