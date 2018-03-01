package healthclub

import "github.com/dpb587/go-schemaorg"

// // A health club.
var Type = schemaorg.NewDataType("http://schema.org", "HealthClub")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
