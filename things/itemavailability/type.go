package itemavailability

import "github.com/dpb587/go-schemaorg"

// // A list of possible product availability options.
var Type = schemaorg.NewDataType("http://schema.org", "ItemAvailability")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
