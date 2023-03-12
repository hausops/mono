package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hausops/mono/apps/dashboard-api/adapter/local"
	"github.com/hausops/mono/apps/dashboard-api/graphql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", newGraphqlServer())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func newGraphqlServer() *handler.Server {
	propertySvc := local.NewPropertyService()
	c := graphql.Config{
		Resolvers: &graphql.Resolver{
			Property: propertySvc,
		},
	}
	return handler.NewDefaultServer(graphql.NewExecutableSchema(c))
}
