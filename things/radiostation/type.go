package radiostation

import "github.com/dpb587/go-schemaorg"

// // A radio station.
var Type = schemaorg.NewDataType("http://schema.org", "RadioStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
