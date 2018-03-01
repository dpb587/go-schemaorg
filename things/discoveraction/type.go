package discoveraction

import "github.com/dpb587/go-schemaorg"

// // The act of discovering/finding an object.
var Type = schemaorg.NewDataType("http://schema.org", "DiscoverAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
