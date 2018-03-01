package episode

import "github.com/dpb587/go-schemaorg"

// // A media episode (e.g. TV, radio, video game) which can be part of a series or
// season.
var Type = schemaorg.NewDataType("http://schema.org", "Episode")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
