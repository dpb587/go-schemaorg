package radioseason

import "github.com/dpb587/go-schemaorg"

// // Season dedicated to radio broadcast and associated online delivery.
var Type = schemaorg.NewDataType("http://schema.org", "RadioSeason")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
