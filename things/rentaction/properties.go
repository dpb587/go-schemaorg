package rentaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of participant. The owner of the real estate property.
	Landlord = schemaorg.NewProperty("landlord")

	// A sub property of participant. The real estate agent involved in the action.
	RealEstateAgent = schemaorg.NewProperty("realEstateAgent")
)
