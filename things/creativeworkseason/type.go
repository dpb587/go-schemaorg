package creativeworkseason

import "github.com/dpb587/go-schemaorg"

// // A media season e.g. tv, radio, video game etc.
var Type = schemaorg.NewDataType("http://schema.org", "CreativeWorkSeason")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
