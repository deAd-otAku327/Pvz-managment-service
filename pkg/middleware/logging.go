package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type contextKey int8

const CtxRequestIDKey contextKey = iota

func Logging(log *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client := r.RemoteAddr

			rid := uuid.New().String()
			w.Header().Set("X-Request-ID", rid)

			log.Info(
				fmt.Sprintf("%s %s", r.Method, r.URL.Path),
				slog.String("client", client),
				slog.String("request_id", rid),
			)

			startReq := time.Now()

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxRequestIDKey, rid)))

			responseTime := time.Since(startReq).Milliseconds()

			log.Info(
				fmt.Sprintf("%s %s", r.Method, r.URL.Path),
				slog.String("client", client),
				slog.String("resp_time (ms)", strconv.Itoa(int(responseTime))),
				slog.String("request_id", rid),
			)
		})
	}
}
