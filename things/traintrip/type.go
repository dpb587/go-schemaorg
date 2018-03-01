package traintrip

import "github.com/dpb587/go-schemaorg"

// // A trip on a commercial train line.
var Type = schemaorg.NewDataType("http://schema.org", "TrainTrip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
