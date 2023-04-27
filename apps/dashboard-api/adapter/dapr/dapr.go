// Package dapr implements domain components that enable communication with
// backend services via Dapr.
package dapr

import (
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Conn creates a gRPC client connection to the Dapr sidecar proxy,
// which enables communication with other Dapr services.
//
// Returns a *grpc.ClientConn if successful, or an error if the port number
// is undefined or if the connection cannot be established.
func Conn() (*grpc.ClientConn, error) {
	port := os.Getenv("DAPR_GRPC_PORT")
	if port == "" {
		return nil, errors.New("DAPR_GRPC_PORT is undefined")
	}

	return grpc.Dial(
		fmt.Sprintf("localhost:%s", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
