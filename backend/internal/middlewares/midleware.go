package middlewares

import (
	"context"
	"net/http"

	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/facades/auth"
	"github.com/mrverdant13/dash_buttons/backend/facades/users"
	"github.com/mrverdant13/dash_buttons/backend/graph/gqlmodel"
)

const (
	ctxName = "user"
)

var userCtxKey = &contextKey{ctxName}

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

				user, err := usersRepo.GetByID(userID)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				ctx := context.WithValue(
					r.Context(),
					userCtxKey,
					user,
				)

				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
			},
		)
	}
}

// CtxUser get the user from the context.
//
// REQUIRES the auth middleware to have run.
func CtxUser(ctx context.Context) *gqlmodel.User {
	user, _ := ctx.Value(userCtxKey).(*gqlmodel.User)
	return user
}

// CtxAdminUser get the admin user from the context.
//
// REQUIRES the auth middleware to have run.
func CtxAdminUser(ctx context.Context) *gqlmodel.User {
	user := CtxUser(ctx)
	if user == nil || !user.IsAdmin {
		return nil
	}
	return user
}
