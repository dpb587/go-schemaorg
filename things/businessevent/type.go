package businessevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Business event.
var Type = schemaorg.NewDataType("http://schema.org", "BusinessEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
