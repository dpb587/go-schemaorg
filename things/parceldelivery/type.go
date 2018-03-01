package parceldelivery

import "github.com/dpb587/go-schemaorg"

// // The delivery of a parcel either via the postal service or a commercial
// service.
var Type = schemaorg.NewDataType("http://schema.org", "ParcelDelivery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
