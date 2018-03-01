package lodgingbusiness

import "github.com/dpb587/go-schemaorg"

// // A lodging business, such as a motel, hotel, or inn.
var Type = schemaorg.NewDataType("http://schema.org", "LodgingBusiness")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
