package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/hausops/mono/services/auth-svc/adapter/redis"
	"github.com/hausops/mono/services/auth-svc/domain/session"
	sessiontesting "github.com/hausops/mono/services/auth-svc/domain/session/testing"
)

func TestSessionRepository(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	srv := miniredis.RunT(t)
	client, err := redis.Conn(ctx, fmt.Sprintf("redis://%s", srv.Addr()))
	if err != nil {
		t.Fatalf("connect to redis: %v", err)
	}

	sessiontesting.TestRepository(t, func(t *testing.T) session.Repository {
		t.Cleanup(func() { srv.FlushAll() })
		return redis.NewSessionRepository(client)
	})
}
