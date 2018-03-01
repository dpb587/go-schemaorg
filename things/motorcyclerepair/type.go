package motorcyclerepair

import "github.com/dpb587/go-schemaorg"

// // A motorcycle repair shop.
var Type = schemaorg.NewDataType("http://schema.org", "MotorcycleRepair")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
