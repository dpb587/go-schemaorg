package apartmentcomplex

import "github.com/dpb587/go-schemaorg"

// // Residence type: Apartment complex.
var Type = schemaorg.NewDataType("http://schema.org", "ApartmentComplex")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
