package dislikeaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a negative sentiment about the object. An agent
// dislikes an object (a proposition, topic or theme) with participants.
var Type = schemaorg.NewDataType("http://schema.org", "DislikeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
