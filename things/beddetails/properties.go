package beddetails

import "github.com/dpb587/go-schemaorg"

var (
	// The type of bed to which the BedDetail refers, i.e. the type of bed available
	// in the quantity indicated by quantity.
	TypeOfBed = schemaorg.NewProperty("typeOfBed")

	// The quantity of the given bed type available in the HotelRoom, Suite, House,
	// or Apartment.
	NumberOfBeds = schemaorg.NewProperty("numberOfBeds")
)
