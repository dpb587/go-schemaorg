package vehicle

import "github.com/dpb587/go-schemaorg"

// // A vehicle is a device that is designed or used to transport people or cargo
// over land, water, air, or through space.
var Type = schemaorg.NewDataType("http://schema.org", "Vehicle")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
