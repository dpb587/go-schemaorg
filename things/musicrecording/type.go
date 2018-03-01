package musicrecording

import "github.com/dpb587/go-schemaorg"

// // A music recording (track), usually a single song.
var Type = schemaorg.NewDataType("http://schema.org", "MusicRecording")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
