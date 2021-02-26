package middlewares

import (
	"context"
	"net/http"

	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/facades/auth"
	"github.com/mrverdant13/dash_buttons/backend/facades/users"
)

const (
	ctxName = "userId"
)

var userIDCtxKey = &contextKey{ctxName}

type contextKey struct {
	name string
}

// Auth is an auth middleware.
func Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				var authService auth.Service
				container.Make(&authService)

				var usersRepo users.Repo
				container.Make(&usersRepo)

				token := r.Header.Get("Authorization")

				if token == "" {
					next.ServeHTTP(w, r)
					return
				}

				userID, err := authService.GetUserIDByToken(token)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				userExists, err := usersRepo.UserWithIDExists(userID)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				ctx := context.WithValue(
					r.Context(),
					userIDCtxKey,
					&userExists,
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
