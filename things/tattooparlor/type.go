package tattooparlor

import "github.com/dpb587/go-schemaorg"

// // A tattoo parlor.
var Type = schemaorg.NewDataType("http://schema.org", "TattooParlor")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
