package museum

import "github.com/dpb587/go-schemaorg"

// // A museum.
var Type = schemaorg.NewDataType("http://schema.org", "Museum")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
