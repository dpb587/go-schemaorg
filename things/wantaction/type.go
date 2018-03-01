package wantaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a desire about the object. An agent wants an object.
var Type = schemaorg.NewDataType("http://schema.org", "WantAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
