package blog

import "github.com/dpb587/go-schemaorg"

// // A blog.
var Type = schemaorg.NewDataType("http://schema.org", "Blog")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
