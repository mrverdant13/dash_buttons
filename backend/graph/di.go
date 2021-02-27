package graph

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/facades/auth"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/facades/districts"
	"github.com/mrverdant13/dash_buttons/backend/facades/provinces"
	"github.com/mrverdant13/dash_buttons/backend/facades/users"
	"github.com/mrverdant13/dash_buttons/backend/graph/generated"
	"github.com/mrverdant13/dash_buttons/backend/internal/middlewares"
)

// Init initializes and runs a GraphQL server.
func Init() {
	var graphQLServerConf config.GraphQLServerConf
	container.Make(&graphQLServerConf)

	graphqlServerPort := graphQLServerConf.PortString()

	var authService auth.Service
	container.Make(&authService)

	var departmentsRepo departments.Repo
	container.Make(&departmentsRepo)

	var provincesRepo provinces.Repo
	container.Make(&provincesRepo)

	var districtsRepo districts.Repo
	container.Make(&districtsRepo)

	var usersRepo users.Repo
	container.Make(&usersRepo)

	router := chi.NewRouter()

	router.Use(middlewares.Auth())

	resolver := NewResolver(
		authService,
		departmentsRepo,
		provincesRepo,
		districtsRepo,
		usersRepo,
	)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver,
			},
		),
	)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground.", graphqlServerPort)
	log.Fatalln(http.ListenAndServe(":"+graphqlServerPort, router).Error())
}
