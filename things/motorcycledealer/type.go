package motorcycledealer

import "github.com/dpb587/go-schemaorg"

// // A motorcycle dealer.
var Type = schemaorg.NewDataType("http://schema.org", "MotorcycleDealer")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
