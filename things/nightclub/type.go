package nightclub

import "github.com/dpb587/go-schemaorg"

// // A nightclub or discotheque.
var Type = schemaorg.NewDataType("http://schema.org", "NightClub")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
