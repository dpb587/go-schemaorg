package busreservation

import "github.com/dpb587/go-schemaorg"

// // A reservation for bus travel. </p>
// 
// <p>Note: This type is for information about actual reservations, e.g. in
// confirmation emails or HTML pages with individual confirmations of
// reservations. For offers of tickets, use <a class="localLink"
// href="http://schema.org/Offer">Offer</a>.
var Type = schemaorg.NewDataType("http://schema.org", "BusReservation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
