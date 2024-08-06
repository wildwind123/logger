package logger

import (
	"context"
	"log/slog"
	"testing"
)

func TestCtxLogger(t *testing.T) {
	t.Skip()
	l := slog.Default()
	ctx := context.Background()
	ctxWithLogger := ToCtx(ctx, l)
	logger := FromCtx(ctxWithLogger)

	logger.Info("test logger")
}
