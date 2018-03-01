package lodgingreservation

import "github.com/dpb587/go-schemaorg"

var (
	// The earliest someone may check into a lodging establishment.
	CheckinTime = schemaorg.NewProperty("checkinTime")

	// The number of children staying in the unit.
	NumChildren = schemaorg.NewProperty("numChildren")

	// A full description of the lodging unit.
	LodgingUnitDescription = schemaorg.NewProperty("lodgingUnitDescription")

	// The latest someone may check out of a lodging establishment.
	CheckoutTime = schemaorg.NewProperty("checkoutTime")

	// Textual description of the unit type (including suite vs. room, size of bed,
	// etc.).
	LodgingUnitType = schemaorg.NewProperty("lodgingUnitType")

	// The number of adults staying in the unit.
	NumAdults = schemaorg.NewProperty("numAdults")
)
