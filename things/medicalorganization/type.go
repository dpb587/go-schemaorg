package medicalorganization

import "github.com/dpb587/go-schemaorg"

// // A medical organization (physical or not), such as hospital, institution or
// clinic.
var Type = schemaorg.NewDataType("http://schema.org", "MedicalOrganization")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
