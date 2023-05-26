package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/hausops/mono/services/user-svc/adapter/mongo"
	"github.com/hausops/mono/services/user-svc/domain/user"
	usertesting "github.com/hausops/mono/services/user-svc/domain/user/testing"
)

func TestUserRepository(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := mongo.Conn(ctx, "mongodb://localhost:27017")
	if err != nil {
		t.Fatalf("connect to mongo: %v", err)
	}
	defer c.Disconnect(ctx)

	uc := c.Database("user-svc-test").Collection("users")
	repo, err := mongo.NewUserRepository(ctx, uc)
	if err != nil {
		t.Fatalf("new user repository (mongo): %v", err)
	}

	usertesting.TestRepository(t, func() (user.Repository, func()) {
		return repo, func() { uc.Drop(ctx) }
	})
}
