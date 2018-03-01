package reservoir

import "github.com/dpb587/go-schemaorg"

// // A reservoir of water, typically an artificially created lake, like the Lake
// Kariba reservoir.
var Type = schemaorg.NewDataType("http://schema.org", "Reservoir")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
