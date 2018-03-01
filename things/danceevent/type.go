package danceevent

import "github.com/dpb587/go-schemaorg"

// // Event type: A social dance.
var Type = schemaorg.NewDataType("http://schema.org", "DanceEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
