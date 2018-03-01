package brewery

import "github.com/dpb587/go-schemaorg"

// // Brewery.
var Type = schemaorg.NewDataType("http://schema.org", "Brewery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
