package readaction

import "github.com/dpb587/go-schemaorg"

// // The act of consuming written content.
var Type = schemaorg.NewDataType("http://schema.org", "ReadAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
