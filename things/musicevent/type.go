package musicevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Music event.
var Type = schemaorg.NewDataType("http://schema.org", "MusicEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
