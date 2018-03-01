package interactaction

import "github.com/dpb587/go-schemaorg"

// // The act of interacting with another person or organization.
var Type = schemaorg.NewDataType("http://schema.org", "InteractAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
