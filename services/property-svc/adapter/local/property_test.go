package local_test

import (
	"testing"

	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/domain/property"

	propertytesting "github.com/hausops/mono/services/property-svc/domain/property/testing"
)

func TestPropertyRepository(t *testing.T) {
	propertytesting.TestRepository(t, func() (property.Repository, func()) {
		return local.NewPropertyRepository(), func() {}
	})
}
