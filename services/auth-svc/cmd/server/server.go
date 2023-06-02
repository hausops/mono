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
	"time"

	"github.com/hausops/mono/services/auth-svc/config"
	"github.com/hausops/mono/services/auth-svc/grpcserver"
	"go.uber.org/zap"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "etc/config.yaml", "path to the config file")

	flag.Parse()

	var conf config.Config
	if err := config.LoadFromFile(configFile, &conf); err != nil {
		log.Fatal(err)
	}

	log.Printf("using config %+v\n", conf)

	logger := newLogger(conf)
	srv, err := grpcserver.New(context.Background(), conf, logger)
	if err != nil {
		log.Fatalf("new grpcserver: %v", err)
	}

	port := os.Getenv("APP_PORT")
	// if port is "" a random free port will be chosen.
	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("server listening on %v", conn.Addr())
		if err := srv.Serve(conn); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for a signal to gracefully shut down the server
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// Call the cancel function to release resources in case of early return
	defer cancel()
	<-ctx.Done()

	timeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeout()
	srv.GracefulStop(timeout)
}

func newLogger(conf config.Config) (logger *zap.Logger) {
	var err error
	switch conf.Mode {
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
