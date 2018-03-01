package howtoitem

import "github.com/dpb587/go-schemaorg"

// // An item used as either a tool or supply when performing the instructions for
// how to to achieve a result.
var Type = schemaorg.NewDataType("http://schema.org", "HowToItem")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
