package slog

import "log/slog"

func NoSpecial() {
	slog.Debug("server started!🚀") // want "log message must not contain special symbols or emojis"
}
