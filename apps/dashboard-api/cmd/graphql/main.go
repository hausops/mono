package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/hausops/mono/apps/dashboard-api/adapter/local"
	"github.com/hausops/mono/apps/dashboard-api/graphql"
	timeout "github.com/vearne/gin-timeout"
)

const defaultPort = "8080"
const graphqlRoot = "/graphql"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(timeout.Timeout(
		timeout.WithTimeout(5*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout),
	))
	r.Use(secure.Secure(secure.Options{
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
	}))

	r.GET("/", playgroundHandler())
	r.POST(graphqlRoot, graphqlHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", graphqlRoot)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler() gin.HandlerFunc {
	propertySvc := local.NewPropertyService()
	cfg := graphql.Config{
		Resolvers: &graphql.Resolver{
			PropertySvc: propertySvc,
		},
	}
	h := handler.NewDefaultServer(graphql.NewExecutableSchema(cfg))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
