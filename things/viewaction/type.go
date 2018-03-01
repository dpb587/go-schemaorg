package viewaction

import "github.com/dpb587/go-schemaorg"

// // The act of consuming static visual content.
var Type = schemaorg.NewDataType("http://schema.org", "ViewAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
