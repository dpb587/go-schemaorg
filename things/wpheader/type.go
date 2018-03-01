package wpheader

import "github.com/dpb587/go-schemaorg"

// // The header section of the page.
var Type = schemaorg.NewDataType("http://schema.org", "WPHeader")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
