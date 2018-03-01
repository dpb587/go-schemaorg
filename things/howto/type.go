package howto

import "github.com/dpb587/go-schemaorg"

// // Instructions that explain how to achieve a result by performing a sequence of
// steps.
var Type = schemaorg.NewDataType("http://schema.org", "HowTo")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
