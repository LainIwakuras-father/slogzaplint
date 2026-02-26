package slog

import "log/slog"

func EnglishCase() {
	slog.Debug("запуск сервера") // want "log message must be in english"
}
