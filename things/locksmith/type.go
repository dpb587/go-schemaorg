package locksmith

import "github.com/dpb587/go-schemaorg"

// // A locksmith.
var Type = schemaorg.NewDataType("http://schema.org", "Locksmith")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
