package automotivebusiness

import "github.com/dpb587/go-schemaorg"

// // Car repair, sales, or parts.
var Type = schemaorg.NewDataType("http://schema.org", "AutomotiveBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
