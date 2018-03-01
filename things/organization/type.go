package organization

import "github.com/dpb587/go-schemaorg"

// // An organization such as a school, NGO, corporation, club, etc.
var Type = schemaorg.NewDataType("http://schema.org", "Organization")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
