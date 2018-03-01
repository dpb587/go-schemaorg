package lodgingreservation

import "github.com/dpb587/go-schemaorg"

// // A reservation for lodging at a hotel, motel, inn, etc.</p>
// 
// <p>Note: This type is for information about actual reservations, e.g. in
// confirmation emails or HTML pages with individual confirmations of
// reservations.
var Type = schemaorg.NewDataType("http://schema.org", "LodgingReservation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
