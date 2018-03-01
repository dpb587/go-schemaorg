package intangible

import "github.com/dpb587/go-schemaorg"

// // A utility class that serves as the umbrella for a number of 'intangible'
// things such as quantities, structured values, etc.
var Type = schemaorg.NewDataType("http://schema.org", "Intangible")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
