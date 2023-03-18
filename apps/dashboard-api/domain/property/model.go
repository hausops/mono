package property

type Property interface {
	IsProperty()
}

type PropertyInfo interface {
	IsPropertyInfo()
	GetID() string
	GetCoverImageURL() *string
	GetAddress() Address
	GetYearBuilt() *int
}

type Address struct {
	Line1 string  `json:"line1"`
	Line2 *string `json:"line2,omitempty"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Zip   string  `json:"zip"`
}

type SingleFamilyProperty struct {
	ID            string                   `json:"id"`
	CoverImageURL *string                  `json:"coverImageUrl,omitempty"`
	Address       Address                  `json:"address"`
	YearBuilt     *int                     `json:"yearBuilt,omitempty"`
	Unit          SingleFamilyPropertyUnit `json:"unit,omitempty"`
}

func (SingleFamilyProperty) IsProperty()     {}
func (SingleFamilyProperty) IsPropertyInfo() {}

func (sp SingleFamilyProperty) GetID() string             { return sp.ID }
func (sp SingleFamilyProperty) GetCoverImageURL() *string { return sp.CoverImageURL }
func (sp SingleFamilyProperty) GetAddress() Address       { return sp.Address }
func (sp SingleFamilyProperty) GetYearBuilt() *int        { return sp.YearBuilt }

type SingleFamilyPropertyUnit struct {
	ID            string         `json:"id"`
	Bedrooms      *float64       `json:"bedrooms,omitempty"`
	Bathrooms     *float64       `json:"bathrooms,omitempty"`
	Size          *float64       `json:"size,omitempty"`
	RentAmount    *float64       `json:"rentAmount,omitempty"`
	ActiveListing *RentalListing `json:"activeListing,omitempty"`
}

type MultiFamilyProperty struct {
	ID            string                    `json:"id"`
	CoverImageURL *string                   `json:"coverImageUrl,omitempty"`
	Address       Address                   `json:"address"`
	YearBuilt     *int                      `json:"yearBuilt,omitempty"`
	Units         []MultiFamilyPropertyUnit `json:"units"`
}

func (MultiFamilyProperty) IsProperty()     {}
func (MultiFamilyProperty) IsPropertyInfo() {}

func (mp MultiFamilyProperty) GetID() string             { return mp.ID }
func (mp MultiFamilyProperty) GetCoverImageURL() *string { return mp.CoverImageURL }
func (mp MultiFamilyProperty) GetAddress() Address       { return mp.Address }
func (mp MultiFamilyProperty) GetYearBuilt() *int        { return mp.YearBuilt }

type MultiFamilyPropertyUnit struct {
	ID            string         `json:"id"`
	Number        string         `json:"number"`
	Bedrooms      *float64       `json:"bedrooms,omitempty"`
	Bathrooms     *float64       `json:"bathrooms,omitempty"`
	Size          *float64       `json:"size,omitempty"`
	RentAmount    *float64       `json:"rentAmount,omitempty"`
	ActiveListing *RentalListing `json:"activeListing,omitempty"`
}

type RentalListing struct {
	ID string `json:"id"`
}
