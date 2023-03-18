package property

type Service interface {
	CreateSingleFamilyProperty(
		CreateSingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	CreateMultiFamilyProperty(
		CreateMultiFamilyPropertyInput,
	) (*MultiFamilyProperty, error)

	FindByID(id string) (Property, error)

	FindAll() ([]Property, error)

	UpdateSingleFamilyPropertyByID(
		id string,
		input UpdateSingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	DeleteByID(id string) (Property, error)
}

type CreateSingleFamilyPropertyInput struct {
	CoverImageURL *string                             `json:"coverImageUrl,omitempty"`
	Address       CreateAddressInput                  `json:"address"`
	BuildYear     *int                                `json:"buildYear,omitempty"`
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
	BuildYear     *int                                 `json:"buildYear,omitempty"`
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
	BuildYear     *int                                 `json:"buildYear,omitempty"`
	Unit          *UpdateSingleFamilyPropertyUnitInput `json:"unit,omitempty"`
}

type UpdateSingleFamilyPropertyUnitInput struct {
	Bedrooms   *float64 `json:"bedrooms,omitempty"`
	Bathrooms  *float64 `json:"bathrooms,omitempty"`
	Size       *float64 `json:"size,omitempty"`
	RentAmount *float64 `json:"rentAmount,omitempty"`
}

type UpdateAdderssInput struct {
	Line1 *string `json:"line1,omitempty"`
	Line2 *string `json:"line2,omitempty"`
	City  *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Zip   *string `json:"zip,omitempty"`
}
