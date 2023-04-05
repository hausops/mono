package main

import (
	"flag"
	"log"
	"net"

	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/grpcserver"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "config/config.prod.yaml", "path to the config file")

	flag.Parse()

	var c config.Config
	if err := config.Load(configFile, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("using config %+v\n", c)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpcserver.New(c)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
