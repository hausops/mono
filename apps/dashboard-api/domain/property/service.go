package property

type Service interface {
	CreateSingleFamilyProperty(
		SingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	CreateMultiFamilyProperty(
		MultiFamilyPropertyInput,
	) (*MultiFamilyProperty, error)

	FindByID(id string) (Property, error)

	FindAll() ([]Property, error)

	DeleteByID(id string) (Property, error)
}

type SingleFamilyPropertyInput struct {
	CoverImageURL *string                       `json:"coverImageUrl"`
	Address       NewAddressInput               `json:"address"`
	BuildYear     *int                          `json:"buildYear"`
	Unit          SingleFamilyPropertyUnitInput `json:"unit"`
}

type SingleFamilyPropertyUnitInput struct {
	Bedrooms   *float64 `json:"bedrooms"`
	Bathrooms  *float64 `json:"bathrooms"`
	Size       *float64 `json:"size"`
	RentAmount *float64 `json:"rentAmount"`
}

type MultiFamilyPropertyInput struct {
	CoverImageURL *string                        `json:"coverImageUrl"`
	Address       NewAddressInput                `json:"address"`
	BuildYear     *int                           `json:"buildYear"`
	Units         []MultiFamilyPropertyUnitInput `json:"units"`
}

type MultiFamilyPropertyUnitInput struct {
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
