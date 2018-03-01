package specialty

import "github.com/dpb587/go-schemaorg"

// // Any branch of a field in which people typically develop specific expertise,
// usually after significant study, time, and effort.
var Type = schemaorg.NewDataType("http://schema.org", "Specialty")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
