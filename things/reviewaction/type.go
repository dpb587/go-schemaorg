package reviewaction

import "github.com/dpb587/go-schemaorg"

// // The act of producing a balanced opinion about the object for an audience. An
// agent reviews an object with participants resulting in a review.
var Type = schemaorg.NewDataType("http://schema.org", "ReviewAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
