package quantitativevalue

import "github.com/dpb587/go-schemaorg"

// // A point value or interval for product characteristics and other purposes.
var Type = schemaorg.NewDataType("http://schema.org", "QuantitativeValue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
