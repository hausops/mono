// Package redis implements domain components backed by redis.
package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Conn creates a redis client with default options and test the connection.
//
// uri is the redis connection string URI.
// Example: "redis://<user>:<pass>@localhost:6379/<db>"
func Conn(ctx context.Context, uri string) (*redis.Client, error) {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = 16
	opt.ReadTimeout = 1 * time.Second
	opt.WriteTimeout = 2 * time.Second

	client := redis.NewClient(opt)

	// Test the connection
	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}
