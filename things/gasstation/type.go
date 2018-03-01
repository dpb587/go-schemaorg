package gasstation

import "github.com/dpb587/go-schemaorg"

// // A gas station.
var Type = schemaorg.NewDataType("http://schema.org", "GasStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
