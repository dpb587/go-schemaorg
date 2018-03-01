package informaction

import "github.com/dpb587/go-schemaorg"

// // The act of notifying someone of information pertinent to them, with no
// expectation of a response.
var Type = schemaorg.NewDataType("http://schema.org", "InformAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
