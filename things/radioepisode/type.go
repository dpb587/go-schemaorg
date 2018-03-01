package radioepisode

import "github.com/dpb587/go-schemaorg"

// // A radio episode which can be part of a series or season.
var Type = schemaorg.NewDataType("http://schema.org", "RadioEpisode")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
