package periodical

import "github.com/dpb587/go-schemaorg"

// // A publication in any medium issued in successive parts bearing numerical or
// chronological designations and intended, such as a magazine, scholarly
// journal, or newspaper to continue indefinitely.</p>
// 
// <p>See also <a
// href="http://blog.schema.org/2014/09/schemaorg-support-for-bibliographic_2.html">blog
// post</a>.
var Type = schemaorg.NewDataType("http://schema.org", "Periodical")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
