package article

import "github.com/dpb587/go-schemaorg"

// // An article, such as a news article or piece of investigative report.
// Newspapers and magazines have articles of many different types and this is
// intended to cover them all.</p>
// 
// <p>See also <a
// href="http://blog.schema.org/2014/09/schemaorg-support-for-bibliographic_2.html">blog
// post</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Article")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
