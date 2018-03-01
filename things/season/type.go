package season

import "github.com/dpb587/go-schemaorg"

// // A media season e.g. tv, radio, video game etc.
var Type = schemaorg.NewDataType("http://schema.org", "Season")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
