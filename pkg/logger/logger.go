package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

const (
	Key = "logger"

	RequesID = "request_id"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("logger.NewLogger: %w", err)
	}

	ctx = context.WithValue(ctx, Key, &Logger{logger: logger})

	return ctx, nil
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequesID) != nil {
		fields = append(fields, zap.String(RequesID, ctx.Value(RequesID).(string)))
	}
	l.logger.Info(msg, fields...)
}
