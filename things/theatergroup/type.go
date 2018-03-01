package theatergroup

import "github.com/dpb587/go-schemaorg"

// // A theater group or company, for example, the Royal Shakespeare Company or
// Druid Theatre.
var Type = schemaorg.NewDataType("http://schema.org", "TheaterGroup")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
