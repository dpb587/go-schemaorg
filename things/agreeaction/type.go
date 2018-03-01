package agreeaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a consistency of opinion with the object. An agent
// agrees to/about an object (a proposition, topic or theme) with participants.
var Type = schemaorg.NewDataType("http://schema.org", "AgreeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
