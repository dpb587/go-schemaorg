package website

import "github.com/dpb587/go-schemaorg"

// // A WebSite is a set of related web pages and other items typically served from
// a single web domain and accessible via URLs.
var Type = schemaorg.NewDataType("http://schema.org", "WebSite")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
