package xmiddleware

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

// adds logger with req_id into context
func NewLoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := ulid.Make().String()

			newLogger := logger.With("req_id", reqID)

			r = r.WithContext(context.WithValue(r.Context(), "logger", newLogger))

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			dur := strconv.FormatInt(time.Since(start).Microseconds(), 10)

			newLogger.Info("Request",
				slog.String("method", r.Method),
				slog.String("url", r.RequestURI),
				slog.String("address", r.RemoteAddr),
				slog.String("duration(ms)", dur),
				slog.Int("code", wrapped.Status()),
			)
		})
	}
}

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}
