package quantity

import "github.com/dpb587/go-schemaorg"

// // Quantities such as distance, time, mass, weight, etc. Particular instances of
// say Mass are entities like '3 Kg' or '4 milligrams'.
var Type = schemaorg.NewDataType("http://schema.org", "Quantity")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
