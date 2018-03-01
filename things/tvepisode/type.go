package tvepisode

import "github.com/dpb587/go-schemaorg"

// // A TV episode which can be part of a series or season.
var Type = schemaorg.NewDataType("http://schema.org", "TVEpisode")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
