package taxireservation

import "github.com/dpb587/go-schemaorg"

var (
	// When a taxi will pickup a passenger or a rental car can be picked up.
	PickupTime = schemaorg.NewProperty("pickupTime")

	// Number of people the reservation should accommodate.
	PartySize = schemaorg.NewProperty("partySize")

	// Where a taxi will pick up a passenger or a rental car can be picked up.
	PickupLocation = schemaorg.NewProperty("pickupLocation")
)
