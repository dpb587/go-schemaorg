package volcano

import "github.com/dpb587/go-schemaorg"

// // A volcano, like Fuji san.
var Type = schemaorg.NewDataType("http://schema.org", "Volcano")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
