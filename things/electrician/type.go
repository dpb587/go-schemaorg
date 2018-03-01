package electrician

import "github.com/dpb587/go-schemaorg"

// // An electrician.
var Type = schemaorg.NewDataType("http://schema.org", "Electrician")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
