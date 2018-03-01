package wpfooter

import "github.com/dpb587/go-schemaorg"

// // The footer section of the page.
var Type = schemaorg.NewDataType("http://schema.org", "WPFooter")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
