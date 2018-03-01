package orderstatus

import "github.com/dpb587/go-schemaorg"

// // Enumerated status values for Order.
var Type = schemaorg.NewDataType("http://schema.org", "OrderStatus")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
