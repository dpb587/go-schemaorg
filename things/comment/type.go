package comment

import "github.com/dpb587/go-schemaorg"

// // A comment on an item - for example, a comment on a blog post. The comment's
// content is expressed via the <a class="localLink"
// href="http://schema.org/text">text</a> property, and its topic via <a
// class="localLink" href="http://schema.org/about">about</a>, properties shared
// with all CreativeWorks.
var Type = schemaorg.NewDataType("http://schema.org", "Comment")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
