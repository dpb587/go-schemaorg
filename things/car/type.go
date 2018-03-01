package car

import "github.com/dpb587/go-schemaorg"

// // A car is a wheeled, self-powered motor vehicle used for transportation.
var Type = schemaorg.NewDataType("http://schema.org", "Car")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
