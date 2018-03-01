package commentaction

import "github.com/dpb587/go-schemaorg"

// // The act of generating a comment about a subject.
var Type = schemaorg.NewDataType("http://schema.org", "CommentAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
