package middleware

import (
	"context"
	"net/http"
	"pvz-service/internal/tokenizer"

	"github.com/gorilla/mux"
)

const (
	CookieName = "token"
)

func AuthOnRoles(t tokenizer.Tokenizer, authRoles map[string]struct{}) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie(CookieName)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			token, err := t.VerifyToken(tokenCookie.Value)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			role, err := token.Claims.GetSubject()
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if _, ok := authRoles[role]; !ok {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserRoleKey, role)))
		})
	}
}
