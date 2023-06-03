package dapr

import (
	"context"
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
func Conn(ctx context.Context) (*grpc.ClientConn, error) {
	port := os.Getenv("DAPR_GRPC_PORT")
	if port == "" {
		return nil, errors.New("DAPR_GRPC_PORT is undefined")
	}

	return grpc.DialContext(ctx,
		fmt.Sprintf("localhost:%s", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
}
