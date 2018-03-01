package returnaction

import "github.com/dpb587/go-schemaorg"

// // The act of returning to the origin that which was previously received
// (concrete objects) or taken (ownership).
var Type = schemaorg.NewDataType("http://schema.org", "ReturnAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
