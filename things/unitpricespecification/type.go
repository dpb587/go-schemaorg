package unitpricespecification

import "github.com/dpb587/go-schemaorg"

// // The price asked for a given offer by the respective organization or person.
var Type = schemaorg.NewDataType("http://schema.org", "UnitPriceSpecification")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
