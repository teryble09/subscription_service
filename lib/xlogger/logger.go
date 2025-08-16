package xlogger

import (
	"context"
	"log/slog"
)

// gets req_id from ctx
func WithReqID(logger *slog.Logger, ctx context.Context) *slog.Logger {
	reqID := ctx.Value("req_id").(string)
	return logger.With("req_id", reqID)
}
