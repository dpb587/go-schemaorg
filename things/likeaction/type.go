package likeaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a positive sentiment about the object. An agent likes
// an object (a proposition, topic or theme) with participants.
var Type = schemaorg.NewDataType("http://schema.org", "LikeAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
