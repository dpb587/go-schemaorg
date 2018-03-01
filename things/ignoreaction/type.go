package ignoreaction

import "github.com/dpb587/go-schemaorg"

// // The act of intentionally disregarding the object. An agent ignores an object.
var Type = schemaorg.NewDataType("http://schema.org", "IgnoreAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
