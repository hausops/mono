package local_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/property-svc/adapter/local"
	"github.com/hausops/mono/services/property-svc/domain/property"
)

func TestPropertyRepository(t *testing.T) {
	ctx := context.Background()

	sample := local.ExampleProperties()

	repo := local.NewPropertyRepository()
	repo.ReplaceProperties(sample)

	t.Run("FindByID", func(t *testing.T) {
		testCases := []struct {
			id   uuid.UUID
			want property.Property
			err  error
		}{
			{
				id:   sample[0].GetID(),
				want: sample[0],
				err:  nil,
			},
			{
				id:   sample[4].GetID(),
				want: sample[4],
				err:  nil,
			},
			{
				id:   uuid.New(),
				want: nil,
				err:  property.ErrNotFound,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.id.String(), func(t *testing.T) {
				got, err := repo.FindByID(ctx, tc.id)

				if err != tc.err {
					t.Fatalf("want error %q but got %q", tc.err, err)
				}

				if got != tc.want {
					t.Fatalf("want %v but got %v", tc.want, got)
				}
			})
		}
	})
}
