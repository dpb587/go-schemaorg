package mosque

import "github.com/dpb587/go-schemaorg"

// // A mosque.
var Type = schemaorg.NewDataType("http://schema.org", "Mosque")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
