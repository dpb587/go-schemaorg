package comedyclub

import "github.com/dpb587/go-schemaorg"

// // A comedy club.
var Type = schemaorg.NewDataType("http://schema.org", "ComedyClub")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
