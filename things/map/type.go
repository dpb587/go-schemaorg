package map

import "github.com/dpb587/go-schemaorg"

// // A map.
var Type = schemaorg.NewDataType("http://schema.org", "Map")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
