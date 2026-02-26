package zap

import "go.uber.org/zap"

func lowerCases() {
	logger := zap.NewNop()
	sugar := logger.Sugar()

	sugar.Infow("Starting server") // want "log message must start with a lowercase letter"
}
