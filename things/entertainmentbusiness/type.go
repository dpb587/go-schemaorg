package entertainmentbusiness

import "github.com/dpb587/go-schemaorg"

// // A business providing entertainment.
var Type = schemaorg.NewDataType("http://schema.org", "EntertainmentBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
