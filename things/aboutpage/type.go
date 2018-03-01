package aboutpage

import "github.com/dpb587/go-schemaorg"

// // Web page type: About page.
var Type = schemaorg.NewDataType("http://schema.org", "AboutPage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
