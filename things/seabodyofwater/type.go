package seabodyofwater

import "github.com/dpb587/go-schemaorg"

// // A sea (for example, the Caspian sea).
var Type = schemaorg.NewDataType("http://schema.org", "SeaBodyOfWater")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
