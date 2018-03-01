package theaterevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Theater performance.
var Type = schemaorg.NewDataType("http://schema.org", "TheaterEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
