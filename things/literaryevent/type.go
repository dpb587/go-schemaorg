package literaryevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Literary event.
var Type = schemaorg.NewDataType("http://schema.org", "LiteraryEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
