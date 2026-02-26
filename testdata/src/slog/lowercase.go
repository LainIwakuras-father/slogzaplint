package slog

import (
	"log/slog"
)

func LowerCase() {
	slog.Info("Starting server") // want "log message must start with a lowercase letter"
	slog.Info("starting server") // OK
}
