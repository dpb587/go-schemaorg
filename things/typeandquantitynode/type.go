package typeandquantitynode

import "github.com/dpb587/go-schemaorg"

// // A structured value indicating the quantity, unit of measurement, and business
// function of goods included in a bundle offer.
var Type = schemaorg.NewDataType("http://schema.org", "TypeAndQuantityNode")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
