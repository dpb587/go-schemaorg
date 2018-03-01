package policestation

import "github.com/dpb587/go-schemaorg"

// // A police station.
var Type = schemaorg.NewDataType("http://schema.org", "PoliceStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
