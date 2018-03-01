package educationevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Education event.
var Type = schemaorg.NewDataType("http://schema.org", "EducationEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
