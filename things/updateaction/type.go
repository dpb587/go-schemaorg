package updateaction

import "github.com/dpb587/go-schemaorg"

// // The act of managing by changing/editing the state of the object.
var Type = schemaorg.NewDataType("http://schema.org", "UpdateAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
