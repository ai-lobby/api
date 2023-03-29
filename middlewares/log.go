package middlewares

import (
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

func LogRequest(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Info("request started", slog.String("req_body", "123"))
		next.ServeHTTP(w, r)
		logger.Info("request responded", slog.Duration("took", time.Since(start)))
	})
}
