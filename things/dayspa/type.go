package dayspa

import "github.com/dpb587/go-schemaorg"

// // A day spa.
var Type = schemaorg.NewDataType("http://schema.org", "DaySpa")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
