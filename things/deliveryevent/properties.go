package deliveryevent

import "github.com/dpb587/go-schemaorg"

var (
	// Method used for delivery or shipping.
	HasDeliveryMethod = schemaorg.NewProperty("hasDeliveryMethod")

	// After this date, the item will no longer be available for pickup.
	AvailableThrough = schemaorg.NewProperty("availableThrough")

	// When the item is available for pickup from the store, locker, etc.
	AvailableFrom = schemaorg.NewProperty("availableFrom")

	// Password, PIN, or access code needed for delivery (e.g. from a locker).
	AccessCode = schemaorg.NewProperty("accessCode")
)
