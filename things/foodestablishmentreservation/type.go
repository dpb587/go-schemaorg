package foodestablishmentreservation

import "github.com/dpb587/go-schemaorg"

// // A reservation to dine at a food-related business.Note: This type is for
// information about actual reservations, e.g. in confirmation emails or HTML
// pages with individual confirmations of reservations.
var Type = schemaorg.NewDataType("http://schema.org", "FoodEstablishmentReservation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
