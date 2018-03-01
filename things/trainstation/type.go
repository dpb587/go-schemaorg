package trainstation

import "github.com/dpb587/go-schemaorg"

// // A train station.
var Type = schemaorg.NewDataType("http://schema.org", "TrainStation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
