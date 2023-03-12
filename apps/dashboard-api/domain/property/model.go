package property

type Property interface {
	IsProperty()
}

type PropertyModel interface {
	IsPropertyModel()
	GetID() string
	GetCoverImageURL() *string
	GetAddress() *Address
	GetBuildYear() *int
}

type Address struct {
	Line1 string  `json:"line1"`
	Line2 *string `json:"line2"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Zip   string  `json:"zip"`
}

type SingleFamilyProperty struct {
	ID            string            `json:"id"`
	CoverImageURL *string           `json:"coverImageUrl"`
	Address       *Address          `json:"address"`
	BuildYear     *int              `json:"buildYear"`
	Unit          *SingleFamilyUnit `json:"unit"`
}

func (SingleFamilyProperty) IsProperty()      {}
func (SingleFamilyProperty) IsPropertyModel() {}

func (sp SingleFamilyProperty) GetID() string             { return sp.ID }
func (sp SingleFamilyProperty) GetCoverImageURL() *string { return sp.CoverImageURL }
func (sp SingleFamilyProperty) GetAddress() *Address      { return sp.Address }
func (sp SingleFamilyProperty) GetBuildYear() *int        { return sp.BuildYear }

type SingleFamilyUnit struct {
	ID            string         `json:"id"`
	Bedrooms      *float64       `json:"bedrooms"`
	Bathrooms     *float64       `json:"bathrooms"`
	Size          *float64       `json:"size"`
	RentAmount    *float64       `json:"rentAmount"`
	ActiveListing *RentalListing `json:"activeListing"`
}

type MultiFamilyProperty struct {
	ID            string                     `json:"id"`
	CoverImageURL *string                    `json:"coverImageUrl"`
	Address       *Address                   `json:"address"`
	BuildYear     *int                       `json:"buildYear"`
	Units         []*MultiFamilyPropertyUnit `json:"units"`
}

func (MultiFamilyProperty) IsProperty()      {}
func (MultiFamilyProperty) IsPropertyModel() {}

func (mp MultiFamilyProperty) GetID() string             { return mp.ID }
func (mp MultiFamilyProperty) GetCoverImageURL() *string { return mp.CoverImageURL }
func (mp MultiFamilyProperty) GetAddress() *Address      { return mp.Address }
func (mp MultiFamilyProperty) GetBuildYear() *int        { return mp.BuildYear }

type MultiFamilyPropertyUnit struct {
	ID            string         `json:"id"`
	Number        string         `json:"number"`
	Bedrooms      *float64       `json:"bedrooms"`
	Bathrooms     *float64       `json:"bathrooms"`
	Size          *float64       `json:"size"`
	RentAmount    *float64       `json:"rentAmount"`
	ActiveListing *RentalListing `json:"activeListing"`
}

type RentalListing struct {
	ID string `json:"id"`
}
