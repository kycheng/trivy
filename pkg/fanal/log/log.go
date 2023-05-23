package log

import (
	"time"

	"github.com/aquasecurity/trivy/pkg/bug"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	defer func(start time.Time) { bug.PrintCustomStack(start) }(time.Now())
	if logger, err := zap.NewProduction(); err == nil {
		Logger = logger.Sugar()
	}
}

func SetLogger(l *zap.SugaredLogger) {
	Logger = l
}
