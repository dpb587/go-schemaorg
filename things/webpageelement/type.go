package webpageelement

import "github.com/dpb587/go-schemaorg"

// // A web page element, like a table or an image.
var Type = schemaorg.NewDataType("http://schema.org", "WebPageElement")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
