package disagreeaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a difference of opinion with the object. An agent
// disagrees to/about an object (a proposition, topic or theme) with
// participants.
var Type = schemaorg.NewDataType("http://schema.org", "DisagreeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
