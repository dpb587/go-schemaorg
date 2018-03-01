package servicechannel

import "github.com/dpb587/go-schemaorg"

// // A means for accessing a service, e.g. a government office location, web site,
// or phone number.
var Type = schemaorg.NewDataType("http://schema.org", "ServiceChannel")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
