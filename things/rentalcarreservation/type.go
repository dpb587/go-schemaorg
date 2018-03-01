package rentalcarreservation

import "github.com/dpb587/go-schemaorg"

// // A reservation for a rental car.</p>
// 
// <p>Note: This type is for information about actual reservations, e.g. in
// confirmation emails or HTML pages with individual confirmations of
// reservations.
var Type = schemaorg.NewDataType("http://schema.org", "RentalCarReservation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
