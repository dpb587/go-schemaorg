package contactpage

import "github.com/dpb587/go-schemaorg"

// // Web page type: Contact page.
var Type = schemaorg.NewDataType("http://schema.org", "ContactPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
