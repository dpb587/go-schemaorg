package winery

import "github.com/dpb587/go-schemaorg"

// // A winery.
var Type = schemaorg.NewDataType("http://schema.org", "Winery")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
