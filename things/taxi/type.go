package taxi

import "github.com/dpb587/go-schemaorg"

// // A taxi.
var Type = schemaorg.NewDataType("http://schema.org", "Taxi")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
