package hairsalon

import "github.com/dpb587/go-schemaorg"

// // A hair salon.
var Type = schemaorg.NewDataType("http://schema.org", "HairSalon")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
