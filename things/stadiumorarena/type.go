package stadiumorarena

import "github.com/dpb587/go-schemaorg"

// // A stadium.
var Type = schemaorg.NewDataType("http://schema.org", "StadiumOrArena")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
