package mongo_test

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/hausops/mono/services/user-svc/adapter/mongo"
	"github.com/hausops/mono/services/user-svc/domain/user"
	usertesting "github.com/hausops/mono/services/user-svc/domain/user/testing"
)

// Run this test manually with -mongoURI "mongo://..."
// e.g. go test ./adapter/mongo -mongoURI "mongodb://localhost:27017"
var uri = flag.String("mongoURI", "", "mongodb connection URI string")

func TestUserRepository(t *testing.T) {
	if *uri == "" {
		t.Skip()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client, err := mongo.Conn(ctx, *uri)
	if err != nil {
		t.Fatalf("connect to mongo: %v", err)
	}
	defer client.Disconnect(ctx)

	userCollection := client.Database("user-svc-test").Collection("users")
	repo, err := mongo.NewUserRepository(ctx, userCollection)
	if err != nil {
		t.Fatalf("new user repository (mongo): %v", err)
	}

	usertesting.TestRepository(t, func() (user.Repository, func()) {
		return repo, func() { userCollection.Drop(ctx) }
	})
}
