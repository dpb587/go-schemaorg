package organizationrole

import "github.com/dpb587/go-schemaorg"

// // A subclass of Role used to describe roles within organizations.
var Type = schemaorg.NewDataType("http://schema.org", "OrganizationRole")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
