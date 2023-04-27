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
	"github.com/hausops/mono/apps/dashboard-api/adapter/dapr"
	"github.com/hausops/mono/apps/dashboard-api/graphql"
	timeout "github.com/vearne/gin-timeout"
	"google.golang.org/grpc"
)

const defaultPort = "9098"
const graphqlRoot = "/graphql"

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	// Create dapr connection to connect to other services.
	daprConn, err := dapr.Conn()
	if err != nil {
		log.Fatalf("cannot connect to dapr sidecar: %v", err)
	}
	defer daprConn.Close()

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
	r.POST(graphqlRoot, graphqlHandler(daprConn))

	addr := "localhost:" + port
	log.Printf("server listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", graphqlRoot)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler(conn *grpc.ClientConn) gin.HandlerFunc {
	propertySvc := dapr.NewPropertyService(conn)
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
