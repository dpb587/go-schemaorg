package canal

import "github.com/dpb587/go-schemaorg"

// // A canal, like the Panama Canal.
var Type = schemaorg.NewDataType("http://schema.org", "Canal")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
