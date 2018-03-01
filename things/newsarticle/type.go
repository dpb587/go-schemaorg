package newsarticle

import "github.com/dpb587/go-schemaorg"

// // A NewsArticle is an article whose content reports news, or provides
// background context and supporting materials for understanding the news.</p>
// 
// <p>A more detailed overview of <a href="/docs/news.html">schema.org News
// markup</a> is also available.
var Type = schemaorg.NewDataType("http://schema.org", "NewsArticle")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
