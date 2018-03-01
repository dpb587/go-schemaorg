package beach

import "github.com/dpb587/go-schemaorg"

// // Beach.
var Type = schemaorg.NewDataType("http://schema.org", "Beach")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
