package main

import (
	"log"
	"net"

	"github.com/hausops/mono/services/property-svc/grpcserver"
)

func main() {
	// flag.Parse()

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpcserver.New()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
