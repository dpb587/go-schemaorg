package howtosupply

import "github.com/dpb587/go-schemaorg"

// // A supply consumed when performing the instructions for how to achieve a
// result.
var Type = schemaorg.NewDataType("http://schema.org", "HowToSupply")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
