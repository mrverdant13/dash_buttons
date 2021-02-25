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
	"github.com/mrverdant13/dash_buttons/backend/graph/generated"
)

// Init initializes and runs a GraphQL server.
func Init() {
	var graphQLServerConf config.GraphQLServerConf
	container.Make(&graphQLServerConf)

	graphqlServerPort := graphQLServerConf.PortString()

	var departmentsRepo departments.Repo
	container.Make(&departmentsRepo)

	var authService auth.Service
	container.Make(&authService)

	router := chi.NewRouter()

	resolver := NewResolver(
		departmentsRepo,
		authService,
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
