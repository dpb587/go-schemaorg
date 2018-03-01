package deliverychargespecification

import "github.com/dpb587/go-schemaorg"

// // The price for the delivery of an offer using a particular delivery method.
var Type = schemaorg.NewDataType("http://schema.org", "DeliveryChargeSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
