package comedyevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Comedy event.
var Type = schemaorg.NewDataType("http://schema.org", "ComedyEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
