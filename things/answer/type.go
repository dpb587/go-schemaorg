package answer

import "github.com/dpb587/go-schemaorg"

// // An answer offered to a question; perhaps correct, perhaps opinionated or
// wrong.
var Type = schemaorg.NewDataType("http://schema.org", "Answer")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
