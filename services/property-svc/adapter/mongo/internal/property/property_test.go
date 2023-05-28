package property_test

import (
	"context"
	"flag"
	"testing"
	"time"

	"github.com/hausops/mono/services/property-svc/adapter/mongo"
	"github.com/hausops/mono/services/property-svc/domain/property"
	propertytesting "github.com/hausops/mono/services/property-svc/domain/property/testing"
)

// Run this test manually with -mongoURI "mongo://..."
// e.g. go test ./adapter/mongo -mongoURI "mongodb://localhost:27017"
var uri = flag.String("mongoURI", "", "mongodb connection URI string")

func TestPropertyRepository(t *testing.T) {
	if *uri == "" {
		t.Skip()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	c, err := mongo.Conn(ctx, *uri)
	if err != nil {
		t.Fatalf("connect to mongo: %v", err)
	}
	defer c.Disconnect(ctx)

	pc := c.Database("property-svc-test").Collection("properties")
	repo, err := mongo.NewPropertyRepository(pc)
	if err != nil {
		t.Fatalf("new property repository (mongo): %v", err)
	}

	propertytesting.TestRepository(t, func() (property.Repository, func()) {
		return repo, func() { pc.Drop(ctx) }
	})
}
