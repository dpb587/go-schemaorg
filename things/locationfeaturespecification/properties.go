package locationfeaturespecification

import "github.com/dpb587/go-schemaorg"

var (
	// The hours during which this service or contact is available.
	HoursAvailable = schemaorg.NewProperty("hoursAvailable")

	// The date when the item becomes valid.
	ValidFrom = schemaorg.NewProperty("validFrom")

	// The date after when the item is not valid. For example the end of an offer,
	// salary period, or a period of opening hours.
	ValidThrough = schemaorg.NewProperty("validThrough")
)
