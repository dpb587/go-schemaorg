package publicationvolume

import "github.com/dpb587/go-schemaorg"

// // <p>A part of a successively published publication such as a periodical or
// multi-volume work, often numbered. It may represent a time span, such as a
// year.</p>
// 
// <pre><code>  &lt;br/&gt;&lt;br/&gt;See also &lt;a
// href="http://blog.schema.org/2014/09/schemaorg-support-for-bibliographic_2.html"&gt;blog
// post&lt;/a&gt;.
// </code></pre>
// 
var Type = schemaorg.NewDataType("http://schema.org", "PublicationVolume")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
