package grpcserver

import (
	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/config"
	"github.com/hausops/mono/services/property-svc/domain/property"
)

type dependencies struct {
	propertySvc *property.Service
}

// newDependencies sets up dependencies for running grpcserver based on c.
func newDependencies(c config.Config) dependencies {
	// TODO: handle c.Proxy once adapter/dapr is implemented.
	propertyRepo := local.
		NewPropertyRepository().
		ReplaceProperties(local.ExampleProperties())

	return dependencies{
		propertySvc: property.NewService(propertyRepo),
	}
}
