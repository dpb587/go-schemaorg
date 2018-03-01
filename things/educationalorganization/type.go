package educationalorganization

import "github.com/dpb587/go-schemaorg"

// // An educational organization.
var Type = schemaorg.NewDataType("http://schema.org", "EducationalOrganization")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
