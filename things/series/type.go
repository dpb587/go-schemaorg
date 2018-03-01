package series

import "github.com/dpb587/go-schemaorg"

// // A Series in schema.org is a group of related items, typically but not
// necessarily of the same kind.
var Type = schemaorg.NewDataType("http://schema.org", "Series")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
