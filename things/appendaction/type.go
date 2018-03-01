package appendaction

import "github.com/dpb587/go-schemaorg"

// // The act of inserting at the end if an ordered collection.
var Type = schemaorg.NewDataType("http://schema.org", "AppendAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
