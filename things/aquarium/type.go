package aquarium

import "github.com/dpb587/go-schemaorg"

// // Aquarium.
var Type = schemaorg.NewDataType("http://schema.org", "Aquarium")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
