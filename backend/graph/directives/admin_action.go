package directives

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mrverdant13/dash_buttons/backend/internal/middlewares"
)

func adminActionDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	adminUser := middlewares.CtxAdminUser(ctx)
	if adminUser == nil {
		err := fmt.Errorf("Access denied")
		log.Println(err.Error())
		return nil, err
	}
	return next(ctx)
}
