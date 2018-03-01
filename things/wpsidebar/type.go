package wpsidebar

import "github.com/dpb587/go-schemaorg"

// // A sidebar section of the page.
var Type = schemaorg.NewDataType("http://schema.org", "WPSideBar")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
