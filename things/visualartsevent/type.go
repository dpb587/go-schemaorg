package visualartsevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Visual arts event.
var Type = schemaorg.NewDataType("http://schema.org", "VisualArtsEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
