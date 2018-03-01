package question

import "github.com/dpb587/go-schemaorg"

// // A specific question - e.g. from a user seeking answers online, or collected
// in a Frequently Asked Questions (FAQ) document.
var Type = schemaorg.NewDataType("http://schema.org", "Question")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
