package slog

import "log/slog"

func Sensitive() {
	apiKey := "okritohrjtiowerjotujoe198y47364817359"
	slog.Info("api_key=" + apiKey) // want "log message must not contain sensitive data"
}
