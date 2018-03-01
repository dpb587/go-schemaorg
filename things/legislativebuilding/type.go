package legislativebuilding

import "github.com/dpb587/go-schemaorg"

// // A legislative building&#x2014;for example, the state capitol.
var Type = schemaorg.NewDataType("http://schema.org", "LegislativeBuilding")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
