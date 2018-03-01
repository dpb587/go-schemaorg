package riverbodyofwater

import "github.com/dpb587/go-schemaorg"

// // A river (for example, the broad majestic Shannon).
var Type = schemaorg.NewDataType("http://schema.org", "RiverBodyOfWater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
