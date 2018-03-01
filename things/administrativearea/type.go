package administrativearea

import "github.com/dpb587/go-schemaorg"

// // A geographical region, typically under the jurisdiction of a particular
// government.
var Type = schemaorg.NewDataType("http://schema.org", "AdministrativeArea")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
