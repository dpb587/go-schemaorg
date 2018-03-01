package blogposting

import "github.com/dpb587/go-schemaorg"

// // A blog post.
var Type = schemaorg.NewDataType("http://schema.org", "BlogPosting")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
