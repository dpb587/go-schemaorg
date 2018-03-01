package sportsevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Sports event.
var Type = schemaorg.NewDataType("http://schema.org", "SportsEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
