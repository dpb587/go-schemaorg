package scholarlyarticle

import "github.com/dpb587/go-schemaorg"

// // A scholarly article.
var Type = schemaorg.NewDataType("http://schema.org", "ScholarlyArticle")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
