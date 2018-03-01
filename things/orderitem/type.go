package orderitem

import "github.com/dpb587/go-schemaorg"

// // An order item is a line of an order. It includes the quantity and shipping
// details of a bought offer.
var Type = schemaorg.NewDataType("http://schema.org", "OrderItem")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
