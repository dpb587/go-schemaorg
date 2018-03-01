package hindutemple

import "github.com/dpb587/go-schemaorg"

// // A Hindu temple.
var Type = schemaorg.NewDataType("http://schema.org", "HinduTemple")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
