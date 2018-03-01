package governmentorganization

import "github.com/dpb587/go-schemaorg"

// // A governmental organization or agency.
var Type = schemaorg.NewDataType("http://schema.org", "GovernmentOrganization")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
