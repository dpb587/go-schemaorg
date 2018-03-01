package plumber

import "github.com/dpb587/go-schemaorg"

// // A plumbing service.
var Type = schemaorg.NewDataType("http://schema.org", "Plumber")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
