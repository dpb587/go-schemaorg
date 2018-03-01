package radioseries

import "github.com/dpb587/go-schemaorg"

// // CreativeWorkSeries dedicated to radio broadcast and associated online
// delivery.
var Type = schemaorg.NewDataType("http://schema.org", "RadioSeries")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
