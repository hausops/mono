// Package mongo implements domain components backed by mongodb.
package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Conn creates a mongodb client with default options and test the connection.
//
// uri is the mongodb connection string URI.
// See: https://www.mongodb.com/docs/manual/reference/connection-string/
func Conn(ctx context.Context, uri string) (*mongo.Client, error) {
	opt := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(16).
		SetConnectTimeout(5 * time.Second)

	c, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = c.Ping(ctx, nil); err != nil {
		c.Disconnect(ctx)
		return nil, err
	}

	return c, nil
}