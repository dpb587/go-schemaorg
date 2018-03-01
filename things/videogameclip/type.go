package videogameclip

import "github.com/dpb587/go-schemaorg"

// // A short segment/part of a video game.
var Type = schemaorg.NewDataType("http://schema.org", "VideoGameClip")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
