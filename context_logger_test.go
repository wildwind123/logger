package logger

import (
	"context"
	"log/slog"
	"testing"
)

func TestCtxLogger(t *testing.T) {
	l := slog.Default()
	ctx := context.Background()
	ctxWithLogger := ToCtx(ctx, l)
	logger := FromCtx(ctxWithLogger)

	logger.Info("test logger")
}

func TestRequestIDCtx(t *testing.T) {
	ctx := context.Background()

	sampleRequestID := "sample_request_id"

	ctx = RequestIDToCtx(ctx, sampleRequestID)

	requestID := RequestIDFromCtx(ctx)

	if requestID != sampleRequestID {
		t.Error("wrong request id")
	}
}
