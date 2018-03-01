package lockerdelivery

import "github.com/dpb587/go-schemaorg"

// // A DeliveryMethod in which an item is made available via locker.
var Type = schemaorg.NewDataType("http://schema.org", "LockerDelivery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
