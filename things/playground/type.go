package playground

import "github.com/dpb587/go-schemaorg"

// // A playground.
var Type = schemaorg.NewDataType("http://schema.org", "Playground")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
