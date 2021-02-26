package auth

import (
	"context"
	"net/http"

	"github.com/golobby/container"
)

const (
	ctxName = "user"
)

var userCtxKey = &contextKey{ctxName}

type contextKey struct {
	name string
}

// Middleware is an auth middleware.
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				var authService Service
				container.Make(&authService)

				token := r.Header.Get("Authorization")

				if token == "" {
					next.ServeHTTP(w, r)
					return
				}

				user, err := authService.GetUserByToken(token)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				ctx := context.WithValue(
					r.Context(),
					userCtxKey,
					&user,
				)

				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
			},
		)
	}
}

// ForContext finds the user from the context.
//
// REQUIRES Middleware to have run.
// func ForContext(ctx context.Context) *User {
// 	raw, _ := ctx.Value(userCtxKey).(*User)
// 	return raw
// }
