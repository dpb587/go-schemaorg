package rentalcarreservation

import "github.com/dpb587/go-schemaorg"

var (
	// When a taxi will pickup a passenger or a rental car can be picked up.
	PickupTime = schemaorg.NewProperty("pickupTime")

	// Where a rental car can be dropped off.
	DropoffLocation = schemaorg.NewProperty("dropoffLocation")

	// When a rental car can be dropped off.
	DropoffTime = schemaorg.NewProperty("dropoffTime")

	// Where a taxi will pick up a passenger or a rental car can be picked up.
	PickupLocation = schemaorg.NewProperty("pickupLocation")
)
