package screeningevent

import "github.com/dpb587/go-schemaorg"

// // A screening of a movie or other video.
var Type = schemaorg.NewDataType("http://schema.org", "ScreeningEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
