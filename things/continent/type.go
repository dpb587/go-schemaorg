package continent

import "github.com/dpb587/go-schemaorg"

// // One of the continents (for example, Europe or Africa).
var Type = schemaorg.NewDataType("http://schema.org", "Continent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
