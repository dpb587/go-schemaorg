package openinghoursspecification

import "github.com/dpb587/go-schemaorg"

var (
	// The date when the item becomes valid.
	ValidFrom = schemaorg.NewProperty("validFrom")

	// The date after when the item is not valid. For example the end of an offer,
	// salary period, or a period of opening hours.
	ValidThrough = schemaorg.NewProperty("validThrough")

	// The opening hour of the place or service on the given day(s) of the week.
	Opens = schemaorg.NewProperty("opens")

	// The closing hour of the place or service on the given day(s) of the week.
	Closes = schemaorg.NewProperty("closes")

	// The day of the week for which these opening hours are valid.
	DayOfWeek = schemaorg.NewProperty("dayOfWeek")
)
