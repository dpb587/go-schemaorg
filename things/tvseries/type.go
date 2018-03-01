package tvseries

import "github.com/dpb587/go-schemaorg"

// // CreativeWorkSeries dedicated to TV broadcast and associated online delivery.
var Type = schemaorg.NewDataType("http://schema.org", "TVSeries")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
