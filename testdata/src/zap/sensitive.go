package zap

import "go.uber.org/zap"

func SensitiveCases() {
	logger := zap.NewNop()
	sugar := logger.Sugar()
	apiKey := "rkfoeijtoui4rjetuhiuh3827594y208"
	sugar.Warnw("api_key=" + apiKey) // want "log message must not contain sensitive data"
}
