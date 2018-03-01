package softwareapplication

import "github.com/dpb587/go-schemaorg"

// // A software application.
var Type = schemaorg.NewDataType("http://schema.org", "SoftwareApplication")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
