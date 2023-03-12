package property

type Service interface {
	CreateSingleFamilyProperty(
		CreateSingleFamilyPropertyInput,
	) (*SingleFamilyProperty, error)

	CreateMultiFamilyProperty(
		CreateMultiFamilyPropertyInput,
	) (*MultiFamilyProperty, error)

	FindAll() ([]Property, error)
}

type CreateSingleFamilyPropertyInput struct {
	CoverImageURL *string                             `json:"coverImageUrl"`
	Address       NewAddressInput                     `json:"address"`
	BuildYear     *int                                `json:"buildYear"`
	Unit          CreateSingleFamilyPropertyUnitInput `json:"unit"`
}

func (in CreateSingleFamilyPropertyInput) ToProperty() SingleFamilyProperty {
	return SingleFamilyProperty{
		CoverImageURL: in.CoverImageURL,
		Address:       in.Address.ToAddress(),
		BuildYear:     in.BuildYear,
		Unit:          in.Unit.ToUnit(),
	}
}

type CreateSingleFamilyPropertyUnitInput struct {
	Bedrooms   *float64 `json:"bedrooms"`
	Bathrooms  *float64 `json:"bathrooms"`
	Size       *float64 `json:"size"`
	RentAmount *float64 `json:"rentAmount"`
}

func (in CreateSingleFamilyPropertyUnitInput) ToUnit() SingleFamilyPropertyUnit {
	return SingleFamilyPropertyUnit{
		Bedrooms:   in.Bedrooms,
		Bathrooms:  in.Bathrooms,
		Size:       in.Size,
		RentAmount: in.RentAmount,
	}
}

type CreateMultiFamilyPropertyInput struct {
	CoverImageURL *string                              `json:"coverImageUrl"`
	Address       NewAddressInput                      `json:"address"`
	BuildYear     *int                                 `json:"buildYear"`
	Units         []CreateMultiFamilyPropertyUnitInput `json:"units"`
}

func (in CreateMultiFamilyPropertyInput) ToProperty() MultiFamilyProperty {
	return MultiFamilyProperty{
		CoverImageURL: in.CoverImageURL,
		Address:       in.Address.ToAddress(),
		BuildYear:     in.BuildYear,
	}
}

type CreateMultiFamilyPropertyUnitInput struct {
	Number     string   `json:"number"`
	Bedrooms   *float64 `json:"bedrooms"`
	Bathrooms  *float64 `json:"bathrooms"`
	Size       *float64 `json:"size"`
	RentAmount *float64 `json:"rentAmount"`
}

func (in CreateMultiFamilyPropertyUnitInput) ToUnit() MultiFamilyPropertyUnit {
	return MultiFamilyPropertyUnit{
		Number:     in.Number,
		Bedrooms:   in.Bedrooms,
		Bathrooms:  in.Bathrooms,
		Size:       in.Size,
		RentAmount: in.RentAmount,
	}
}

type NewAddressInput struct {
	Line1 string  `json:"line1"`
	Line2 *string `json:"line2"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Zip   string  `json:"zip"`
}

func (in NewAddressInput) ToAddress() Address {
	return Address(in)
}
