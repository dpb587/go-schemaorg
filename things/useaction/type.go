package useaction

import "github.com/dpb587/go-schemaorg"

// // The act of applying an object to its intended purpose.
var Type = schemaorg.NewDataType("http://schema.org", "UseAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
