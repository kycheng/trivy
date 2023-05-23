package log

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now());
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
}

func SetLogger(l *zap.SugaredLogger) {
	Logger = l
}
