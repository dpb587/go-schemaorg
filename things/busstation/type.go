package busstation

import "github.com/dpb587/go-schemaorg"

// // A bus station.
var Type = schemaorg.NewDataType("http://schema.org", "BusStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
