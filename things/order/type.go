package order

import "github.com/dpb587/go-schemaorg"

// // An order is a confirmation of a transaction (a receipt), which can contain
// multiple line items, each represented by an Offer that has been accepted by
// the customer.
var Type = schemaorg.NewDataType("http://schema.org", "Order")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
