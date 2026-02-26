package zap

import "go.uber.org/zap"

func NoSpecialCases() {
	logger := zap.NewNop()
	sugar := logger.Sugar()

	sugar.Debugw("server started!🚀") // want "log message must not contain special symbols or emojis"
}
