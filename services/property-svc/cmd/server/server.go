package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/grpcserver"
	"go.uber.org/zap"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "config/config.prod.yaml", "path to the config file")

	flag.Parse()

	var c config.Config
	if err := config.LoadByFilename(configFile, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("using config %+v\n", c)

	logger := newLogger(c)
	s := grpcserver.New(c, logger)

	port := os.Getenv("APP_PORT")
	// if port is "" a random free port will be chosen.
	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("server listening on %v", conn.Addr())
		if err := s.Serve(conn); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for a signal to gracefully shut down the server
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// Call the cancel function to release resources in case of early return
	defer cancel()

	<-ctx.Done()
	log.Println("Stopping server gracefully...")
	// TODO: Implement a timeout for shutting down if it takes too long
	s.GracefulStop()
	log.Println("Server shut down successfully")
}

func newLogger(c config.Config) (logger *zap.Logger) {
	var err error
	switch c.Mode {
	case config.ModeDev:
		logger, err = zap.NewDevelopment()
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}
	return logger
}
