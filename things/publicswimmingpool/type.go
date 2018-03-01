package publicswimmingpool

import "github.com/dpb587/go-schemaorg"

// // A public swimming pool.
var Type = schemaorg.NewDataType("http://schema.org", "PublicSwimmingPool")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
