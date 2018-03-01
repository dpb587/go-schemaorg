package televisionstation

import "github.com/dpb587/go-schemaorg"

// // A television station.
var Type = schemaorg.NewDataType("http://schema.org", "TelevisionStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
