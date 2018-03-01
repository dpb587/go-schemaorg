package individualproduct

import "github.com/dpb587/go-schemaorg"

// // A single, identifiable product instance (e.g. a laptop with a particular
// serial number).
var Type = schemaorg.NewDataType("http://schema.org", "IndividualProduct")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
