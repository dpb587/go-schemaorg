package authorizeaction

import "github.com/dpb587/go-schemaorg"

// // The act of granting permission to an object.
var Type = schemaorg.NewDataType("http://schema.org", "AuthorizeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
