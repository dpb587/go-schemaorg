package autorepair

import "github.com/dpb587/go-schemaorg"

// // Car repair business.
var Type = schemaorg.NewDataType("http://schema.org", "AutoRepair")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
