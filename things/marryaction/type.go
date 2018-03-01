package marryaction

import "github.com/dpb587/go-schemaorg"

// // The act of marrying a person.
var Type = schemaorg.NewDataType("http://schema.org", "MarryAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
