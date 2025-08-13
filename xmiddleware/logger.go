package xmiddleware

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

// adds request id in the context for logs with a key "req_id" and logs the details of the request
func NewLoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := ulid.Make().String()
			r = r.WithContext(context.WithValue(r.Context(), "req_id", reqID))

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			dur := strconv.FormatInt(time.Since(start).Microseconds(), 10)

			logger.Info("Request",
				slog.String("req_id", reqID),
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
