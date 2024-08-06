package logger

import (
	"context"
	"log/slog"
)

type ctxKey string

const loggerKey = ctxKey("logger")
const requestIDKey = ctxKey("request_id")

func ToCtx(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromCtx(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		slog.Error("context does not have logger.  will return slog default logger")
		return slog.Default()
	}

	return logger
}

func RequestIDToCtx(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func RequestIDFromCtx(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}
	return requestID
}
