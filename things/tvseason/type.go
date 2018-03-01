package tvseason

import "github.com/dpb587/go-schemaorg"

// // Season dedicated to TV broadcast and associated online delivery.
var Type = schemaorg.NewDataType("http://schema.org", "TVSeason")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
