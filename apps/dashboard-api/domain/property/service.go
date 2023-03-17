package property

type Service interface {
	CreateSingleFamilyProperty(
		CreateSingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	CreateMultiFamilyProperty(
		CreateMultiFamilyPropertyInput,
	) (*MultiFamilyProperty, error)

	FindAll() ([]Property, error)

	DeleteByID(id string) (Property, error)
}

type CreateSingleFamilyPropertyInput struct {
	CoverImageURL *string                             `json:"coverImageUrl"`
	Address       NewAddressInput                     `json:"address"`
	BuildYear     *int                                `json:"buildYear"`
	Unit          CreateSingleFamilyPropertyUnitInput `json:"unit"`
}

type CreateSingleFamilyPropertyUnitInput struct {
	Bedrooms   *float64 `json:"bedrooms"`
	Bathrooms  *float64 `json:"bathrooms"`
	Size       *float64 `json:"size"`
	RentAmount *float64 `json:"rentAmount"`
}

type CreateMultiFamilyPropertyInput struct {
	CoverImageURL *string                              `json:"coverImageUrl"`
	Address       NewAddressInput                      `json:"address"`
	BuildYear     *int                                 `json:"buildYear"`
	Units         []CreateMultiFamilyPropertyUnitInput `json:"units"`
}

type CreateMultiFamilyPropertyUnitInput struct {
	Number     string   `json:"number"`
	Bedrooms   *float64 `json:"bedrooms"`
	Bathrooms  *float64 `json:"bathrooms"`
	Size       *float64 `json:"size"`
	RentAmount *float64 `json:"rentAmount"`
}

type NewAddressInput struct {
	Line1 string  `json:"line1"`
	Line2 *string `json:"line2"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Zip   string  `json:"zip"`
}
