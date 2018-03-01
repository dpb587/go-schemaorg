package deliveryevent

import "github.com/dpb587/go-schemaorg"

// // An event involving the delivery of an item.
var Type = schemaorg.NewDataType("http://schema.org", "DeliveryEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
