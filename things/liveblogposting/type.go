package liveblogposting

import "github.com/dpb587/go-schemaorg"

// // A blog post intended to provide a rolling textual coverage of an ongoing
// event through continuous updates.
var Type = schemaorg.NewDataType("http://schema.org", "LiveBlogPosting")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
