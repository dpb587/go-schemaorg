package animalshelter

import "github.com/dpb587/go-schemaorg"

// // Animal shelter.
var Type = schemaorg.NewDataType("http://schema.org", "AnimalShelter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
