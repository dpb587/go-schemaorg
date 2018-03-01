package publicationissue

import "github.com/dpb587/go-schemaorg"

// // A part of a successively published publication such as a periodical or
// publication volume, often numbered, usually containing a grouping of works
// such as articles.</p>
// 
// <p><a
// href="http://blog.schema.org/2014/09/schemaorg-support-for-bibliographic_2.html">blog
// post</a>.
var Type = schemaorg.NewDataType("http://schema.org", "PublicationIssue")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
