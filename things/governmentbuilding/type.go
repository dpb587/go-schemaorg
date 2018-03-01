package governmentbuilding

import "github.com/dpb587/go-schemaorg"

// // A government building.
var Type = schemaorg.NewDataType("http://schema.org", "GovernmentBuilding")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
