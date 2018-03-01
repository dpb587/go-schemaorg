package searchresultspage

import "github.com/dpb587/go-schemaorg"

// // Web page type: Search results page.
var Type = schemaorg.NewDataType("http://schema.org", "SearchResultsPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
