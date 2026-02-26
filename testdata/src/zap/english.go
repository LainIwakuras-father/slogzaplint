package zap

import "go.uber.org/zap"

func EnglishCases() {
	logger := zap.NewNop()
	sugar := logger.Sugar()

	sugar.Debugw("запуск сервера") // want "log message must be in english"
}
