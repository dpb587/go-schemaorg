package wpadblock

import "github.com/dpb587/go-schemaorg"

// // An advertising section of the page.
var Type = schemaorg.NewDataType("http://schema.org", "WPAdBlock")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
