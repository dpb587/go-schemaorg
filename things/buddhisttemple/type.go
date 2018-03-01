package buddhisttemple

import "github.com/dpb587/go-schemaorg"

// // A Buddhist temple.
var Type = schemaorg.NewDataType("http://schema.org", "BuddhistTemple")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
