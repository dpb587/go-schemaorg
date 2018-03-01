package educationalaudience

import "github.com/dpb587/go-schemaorg"

// // An EducationalAudience.
var Type = schemaorg.NewDataType("http://schema.org", "EducationalAudience")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
