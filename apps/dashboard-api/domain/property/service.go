package property

import "context"

type Service interface {
	CreateSingleFamilyProperty(
		ctx context.Context,
		in CreateSingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	CreateMultiFamilyProperty(
		ctx context.Context,
		in CreateMultiFamilyPropertyInput,
	) (*MultiFamilyProperty, error)

	FindByID(ctx context.Context, id string) (Property, error)

	FindAll(ctx context.Context) ([]Property, error)

	UpdateSingleFamilyPropertyByID(
		ctx context.Context,
		id string,
		in UpdateSingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	UpdateMultiFamilyPropertyByID(
		ctx context.Context,
		id string,
		in UpdateMultiFamilyPropertyInput,
	) (*MultiFamilyProperty, error)

	DeleteByID(ctx context.Context, id string) (Property, error)
}

type CreateSingleFamilyPropertyInput struct {
	CoverImageURL *string                             `json:"coverImageUrl,omitempty"`
	Address       CreateAddressInput                  `json:"address"`
	YearBuilt     *int                                `json:"yearBuilt,omitempty"`
	Unit          CreateSingleFamilyPropertyUnitInput `json:"unit"`
}

type CreateSingleFamilyPropertyUnitInput struct {
	Bedrooms   *float64 `json:"bedrooms,omitempty"`
	Bathrooms  *float64 `json:"bathrooms,omitempty"`
	Size       *float64 `json:"size,omitempty"`
	RentAmount *float64 `json:"rentAmount,omitempty"`
}

type CreateMultiFamilyPropertyInput struct {
	CoverImageURL *string                              `json:"coverImageUrl,omitempty"`
	Address       CreateAddressInput                   `json:"address"`
	YearBuilt     *int                                 `json:"yearBuilt,omitempty"`
	Units         []CreateMultiFamilyPropertyUnitInput `json:"units"`
}

type CreateMultiFamilyPropertyUnitInput struct {
	Number     string   `json:"number"`
	Bedrooms   *float64 `json:"bedrooms"`
	Bathrooms  *float64 `json:"bathrooms"`
	Size       *float64 `json:"size"`
	RentAmount *float64 `json:"rentAmount"`
}

type CreateAddressInput struct {
	Line1 string  `json:"line1"`
	Line2 *string `json:"line2,omitempty"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Zip   string  `json:"zip"`
}

type UpdateSingleFamilyPropertyInput struct {
	CoverImageURL *string                              `json:"coverImageUrl,omitempty"`
	Address       *UpdateAdderssInput                  `json:"address,omitempty"`
	YearBuilt     *int                                 `json:"yearBuilt,omitempty"`
	Unit          *UpdateSingleFamilyPropertyUnitInput `json:"unit,omitempty"`
}

type UpdateSingleFamilyPropertyUnitInput struct {
	Bedrooms   *float64 `json:"bedrooms,omitempty"`
	Bathrooms  *float64 `json:"bathrooms,omitempty"`
	Size       *float64 `json:"size,omitempty"`
	RentAmount *float64 `json:"rentAmount,omitempty"`
}

type UpdateMultiFamilyPropertyInput struct {
	CoverImageURL *string             `json:"coverImageUrl"`
	Address       *UpdateAdderssInput `json:"address"`
	YearBuilt     *int                `json:"yearBuilt"`
}

type UpdateAdderssInput struct {
	Line1 *string `json:"line1,omitempty"`
	Line2 *string `json:"line2,omitempty"`
	City  *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Zip   *string `json:"zip,omitempty"`
}
