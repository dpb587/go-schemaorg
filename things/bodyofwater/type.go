package bodyofwater

import "github.com/dpb587/go-schemaorg"

// // A body of water, such as a sea, ocean, or lake.
var Type = schemaorg.NewDataType("http://schema.org", "BodyOfWater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
