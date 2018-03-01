package videogame

import "github.com/dpb587/go-schemaorg"

// // A video game is an electronic game that involves human interaction with a
// user interface to generate visual feedback on a video device.
var Type = schemaorg.NewDataType("http://schema.org", "VideoGame")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
