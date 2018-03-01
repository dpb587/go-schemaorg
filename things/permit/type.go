package permit

import "github.com/dpb587/go-schemaorg"

// // A permit issued by an organization, e.g. a parking pass.
var Type = schemaorg.NewDataType("http://schema.org", "Permit")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
