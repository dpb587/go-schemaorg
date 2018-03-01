package civicstructure

import "github.com/dpb587/go-schemaorg"

// // A public structure, such as a town hall or concert hall.
var Type = schemaorg.NewDataType("http://schema.org", "CivicStructure")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
