package employeerole

import "github.com/dpb587/go-schemaorg"

// // A subclass of OrganizationRole used to describe employee relationships.
var Type = schemaorg.NewDataType("http://schema.org", "EmployeeRole")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
