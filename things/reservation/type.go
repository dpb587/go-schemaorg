package reservation

import "github.com/dpb587/go-schemaorg"

// // Describes a reservation for travel, dining or an event. Some reservations
// require tickets. </p>
// 
// <p>Note: This type is for information about actual reservations, e.g. in
// confirmation emails or HTML pages with individual confirmations of
// reservations. For offers of tickets, restaurant reservations, flights, or
// rental cars, use <a class="localLink"
// href="http://schema.org/Offer">Offer</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Reservation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
