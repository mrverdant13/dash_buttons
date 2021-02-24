package graph

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golobby/container"
	"github.com/mrverdant13/dash_buttons/backend/config"
	"github.com/mrverdant13/dash_buttons/backend/facades/departments"
	"github.com/mrverdant13/dash_buttons/backend/graph/generated"
)

// Init initializes and runs a GraphQL server.
func Init() {

	var departmentsRepo departments.Repo
	container.Make(&departmentsRepo)

	resolver := NewResolver(departmentsRepo)
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver,
			},
		),
	)

	var graphQLServerConf config.GraphQLServerConf
	container.Make(&graphQLServerConf)

	graphqlServerPort := graphQLServerConf.PortString()

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground.", graphqlServerPort)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":"+graphqlServerPort, nil).Error())

}
