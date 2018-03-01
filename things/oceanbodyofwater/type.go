package oceanbodyofwater

import "github.com/dpb587/go-schemaorg"

// // An ocean (for example, the Pacific).
var Type = schemaorg.NewDataType("http://schema.org", "OceanBodyOfWater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
