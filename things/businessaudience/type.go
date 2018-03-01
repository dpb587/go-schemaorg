package businessaudience

import "github.com/dpb587/go-schemaorg"

// // A set of characteristics belonging to businesses, e.g. who compose an item's
// target audience.
var Type = schemaorg.NewDataType("http://schema.org", "BusinessAudience")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
