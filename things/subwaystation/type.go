package subwaystation

import "github.com/dpb587/go-schemaorg"

// // A subway station.
var Type = schemaorg.NewDataType("http://schema.org", "SubwayStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
